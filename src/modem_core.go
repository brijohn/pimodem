package main

import (
	"fmt"
	"time"
)

const (
	InactivityTimer = iota
	GuardTimer
)

type Modem struct {
	serialChannel chan []byte

	clock *Clock

	guardExpired bool
	breakCount   int

	serial *SerialPort
	line   *Line

	ATHeader    [2]byte
	Command     []byte
	LastCommand []byte

	reg *Registers

	CommandMode bool
}

func (mdm *Modem) readRegister(reg SRegister) byte {
	val, err := mdm.reg.Read(reg)
	if err != nil {
		panic(err)
	}
	return val
}

func (mdm *Modem) writeRegister(reg SRegister, value byte) error {
	err := mdm.reg.Write(reg, value)
	if err != nil {
		panic(err)
	}
	return nil
}

func (mdm *Modem) readSerial() {
	for {
		bytes := make([]byte, 256)
		read, err := mdm.serial.Read(bytes)
		if err == nil {
			mdm.serialChannel <- bytes[:read]
		}
	}
}

func (mdm *Modem) sendCRLF() {
	cr := mdm.readRegister(RegCarriageReturnChar)
	lf := mdm.readRegister(RegLineFeedChar)
	mdm.serial.Write([]byte{cr, lf})
}

func (mdm *Modem) sendResponse(e error) {
	r, ok := e.(*ModemResponseError)
	if ok {
		options := mdm.readRegister(RegStatusOptions)
		if options&Quiet != Quiet {
			if options&Verbose == Verbose {
				mdm.serial.Write([]byte(r.Response()))
			} else {
				mdm.serial.Write([]byte(fmt.Sprintf("%d", r.code)))
			}
			mdm.sendCRLF()
		}
	}
}

func (mdm *Modem) setDataMode(enable bool) {
	if enable {
		mdm.CommandMode = false
		mdm.line.Pause(false)
		timeout := time.Duration(mdm.readRegister(RegInactivityTimeout))
		guardTime := time.Duration(mdm.readRegister(RegEscapeSeqGuardTime))
		mdm.clock.SetDuration(InactivityTimer, time.Second*(timeout*10))
		mdm.clock.SetDuration(GuardTimer, time.Millisecond*(guardTime*20))
		fmt.Println("Entering Datamode")
	} else {
		mdm.CommandMode = true
		mdm.line.Pause(true)
		mdm.clock.SetDuration(GuardTimer, 0)
		mdm.clock.SetDuration(InactivityTimer, 0)
		fmt.Println("Leaving Datamode")
	}
}

func (mdm *Modem) resetInactivityTimer() {
	timeout := time.Duration(mdm.readRegister(RegInactivityTimeout))
	mdm.clock.SetDuration(InactivityTimer, time.Second*(timeout*10))

}

func (mdm *Modem) resetGuardTimer() {
	guardTime := time.Duration(mdm.readRegister(RegEscapeSeqGuardTime))
	mdm.clock.SetDuration(GuardTimer, time.Millisecond*(guardTime*20))
	mdm.guardExpired = false
}

func NewModem(dev string, baud uint32, address string) (*Modem, error) {
	var mdm Modem
	var err error
	mdm.serialChannel = make(chan []byte)
	mdm.CommandMode = true
	mdm.reg = NewRegisters()
	mdm.serial, err = OpenSerial(dev, baud)
	if err != nil {
		return nil, err
	}
	mdm.line, err = NewLine(address)
	if err != nil {
		mdm.serial.Close()
		return nil, err
	}
	mdm.breakCount = 0
	mdm.guardExpired = false
	mdm.clock = NewClock(time.Nanosecond)
	return &mdm, nil
}

func (mdm *Modem) handleCommandInput(input []byte) {
	for _, char := range input {
		if mdm.readRegister(RegStatusOptions)&Echo == Echo {
			mdm.serial.Write([]byte{char})
		}
		if string(mdm.ATHeader[:]) == "AT" || string(mdm.ATHeader[:]) == "at" {
			if char == mdm.readRegister(RegCarriageReturnChar) {
				fmt.Printf("Evalulating %s%s\n", mdm.ATHeader, mdm.Command)
				mdm.LastCommand = append([]byte{}, mdm.Command...)
				handler, err := mdm.Parse()
				for handler != nil {
					handler, err = handler(mdm)
				}
				mdm.sendCRLF()
				mdm.sendResponse(err)
				fmt.Printf("%s\n", err.Error())
				mdm.Command = make([]byte, 0)
				mdm.ATHeader = [2]byte{0, 0}
			} else if char == mdm.readRegister(RegBackspaceChar) {
				if len(mdm.Command) > 0 {
					mdm.Command = mdm.Command[:len(mdm.Command)-1]
					if mdm.readRegister(RegStatusOptions)&Echo == Echo {
						mdm.serial.Write([]byte{' ', mdm.readRegister(RegBackspaceChar)})
					}
				} else {
					if mdm.readRegister(RegStatusOptions)&Echo == Echo {
						mdm.serial.Write(mdm.ATHeader[1:])
					}
				}
			} else {
				mdm.Command = append(mdm.Command, char)
			}
		} else {
			switch {
			case mdm.ATHeader[0] == 'A' && char == 'T':
				fallthrough
			case mdm.ATHeader[0] == 'a' && char == 't':
				mdm.ATHeader[1] = char
			case mdm.ATHeader[0] == 'A' && char == '/':
				fallthrough
			case mdm.ATHeader[0] == 'a' && char == '/':
				mdm.Command = append([]byte{}, mdm.LastCommand...)
				handler, err := mdm.Parse()
				for handler != nil {
					handler, err = handler(mdm)
				}
				mdm.sendCRLF()
				mdm.sendResponse(err)
				fmt.Printf("%s\n", err.Error())
				mdm.Command = make([]byte, 0)
				mdm.ATHeader = [2]byte{0, 0}
			default:
				mdm.ATHeader[0] = char
			}
		}
	}

}

func (mdm *Modem) Start() {
	go mdm.readSerial()
	fmt.Println("Starting Modem")
	for {
		select {
		case <-mdm.clock.GetTimer(InactivityTimer):
			mdm.line.Hangup()
			mdm.setDataMode(false)
		case <-mdm.clock.GetTimer(GuardTimer):
			if mdm.breakCount == 3 {
				mdm.setDataMode(false)
			}
			mdm.guardExpired = true
			mdm.breakCount = 0
		case bytes := <-mdm.serialChannel:
			if mdm.CommandMode {
				mdm.handleCommandInput(bytes)
			} else {
				mdm.resetInactivityTimer()
				mdm.line.Write(bytes)
				if mdm.guardExpired || (mdm.breakCount > 0) {
					for _, ch := range bytes {
						if ch != mdm.readRegister(RegEscapeSeqChar) {
							mdm.breakCount = 0
							break
						}
						mdm.breakCount++
					}
				}
				mdm.resetGuardTimer()
			}
			fmt.Printf("%d bytes recieved %x\n", len(bytes), bytes)
		case bytes := <-mdm.line.Data:
			fmt.Printf("%d bytes recieved from remote %x\n", len(bytes), bytes)
			mdm.resetInactivityTimer()
			mdm.serial.Write(bytes)
		case response := <-mdm.line.Response:
			mdm.sendResponse(NewResponse(response, ""))
			if response == Connect {
				mdm.setDataMode(true)
			} else if response == NoCarrier {
				mdm.setDataMode(false)
			} else if response == Ring {
				autoAnswer := mdm.readRegister(RegAutoAnswer)
				if autoAnswer > 0 {
					ringCount := mdm.readRegister(RegRingCount) + 1
					if autoAnswer == ringCount {
						ringCount = 0
						mdm.line.Pickup()
					}
					mdm.writeRegister(RegRingCount, ringCount)
				}
			}
		}
	}
}
