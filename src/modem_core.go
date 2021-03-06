/** modem_core.go
 * Copyright (C) 2018-2020  Brian Johnson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
	"fmt"
	"github.com/brijohn/pimodem/nvmem"
	"strings"
	"syscall"
	"time"
)

const (
	InactivityTimer = iota
	GuardTimer
	DTRTimer
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

	nvmem *nvmem.NVMEM
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

func (mdm *Modem) autoAnswerRegSetCallback(value byte) {
	if value == 0 {
		DeassertPin("aa")
	} else {
		AssertPin("aa")
	}
}

func (mdm *Modem) bitmapOptionsRegSetCallback(value byte) {
	dcd := mdm.readRegister(RegGeneralBitmapOptions) & 0x20
	dsr := mdm.readRegister(RegGeneralBitmapOptions) & 0x40
	if dcd == 0 {
		AssertPin("dcd")
	} else {
		if !mdm.line.Established() {
			DeassertPin("dcd")
		} else {
			AssertPin("dcd")
		}
	}
	if dsr == 0 {
		AssertPin("dsr")
	} else {
		if !mdm.line.Established() {
			DeassertPin("dsr")
		} else {
			AssertPin("dsr")
		}
	}
}

func (mdm *Modem) setVolume(volume int) {
	logger.Info().Int("volume", volume).Msg("Set modem volume")
	val := mdm.readRegister(RegSpeakerResultsOptions) & VolumeMask
	switch volume {
	case 0:
	case 1:
		val |= VolumeLow
	case 2:
		val |= VolumeMed
	case 3:
		val |= VolumeHigh
	default:
		logger.Debug().Int("volume", volume).Msg("Invalid volume level")
	}
	mdm.writeRegister(RegSpeakerResultsOptions, val)
	mdm.line.Volume = int(val & 0x3)
}

func (mdm *Modem) getVolume() int {
	return int(mdm.readRegister(RegSpeakerResultsOptions) &^ VolumeMask)
}

func (mdm *Modem) setSpeakers(speaker int) {
	logger.Info().Int("speaker", speaker).Msg("Set modem speaker control")
	val := mdm.readRegister(RegSpeakerResultsOptions) & SpeakerMask
	switch speaker {
	case 0:
	case 1:
		val |= SpeakerOffCarrier
	case 2:
		val |= SpeakerOn
	case 3:
		val |= SpeakerOnHandshake
	default:
		logger.Debug().Int("speaker", speaker).Msg("Invalid speaker control value")
	}
	mdm.writeRegister(RegSpeakerResultsOptions, val)
	mdm.line.Speaker = ((val >> 2) & 0x3) != 0
}

func (mdm *Modem) getSpeakers() int {
	return int((mdm.readRegister(RegSpeakerResultsOptions) &^ SpeakerMask) >> 2)
}

func (mdm *Modem) setExtendedResults(results int) {
	logger.Info().Int("rseults", results).Msg("Set extended results")
	val := mdm.readRegister(RegSpeakerResultsOptions) & ResultsLevelMask
	switch results {
	case 0:
	case 1:
		val |= ResultsLevel1
	case 2:
		val |= ResultsLevel2
	case 3:
		val |= ResultsLevel3
	case 4:
		val |= ResultsLevel4
	default:
		logger.Debug().Int("results", results).Msg("Invalid results level")
	}
	mdm.writeRegister(RegSpeakerResultsOptions, val)
}

func (mdm *Modem) getExtendedResultsLevel() int {
	return int((mdm.readRegister(RegSpeakerResultsOptions) &^ ResultsLevelMask))
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
		level := mdm.getExtendedResultsLevel()
		switch r.code {
		case Busy:
			if level != ResultsLevel4 && level != ResultsLevel3 {
				r.code = NoCarrier
			}
		case NoDialtone:
			if level != ResultsLevel4 && level != ResultsLevel2 {
				r.code = NoCarrier
			}
		case Connect:
			if level != 0 {
				r = NewConnectResponseFromSpeed(38400).(*ModemResponseError)
			}
		}
		options := mdm.readRegister(RegStatusOptions)
		if options&Quiet != Quiet {
			if options&Verbose == Verbose {
				mdm.serial.Write([]byte(r.Response()))
			} else {
				mdm.serial.Write([]byte(fmt.Sprintf("%d", r.code)))
			}
			mdm.sendCRLF()
		}
		logger.Info().Err(r).Uint("code", uint(r.code)).Str("response", r.Response()).Msg("Modem Response")
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
		logger.Debug().Msg("Entering Datamode")
	} else {
		mdm.CommandMode = true
		mdm.line.Pause(true)
		mdm.clock.SetDuration(GuardTimer, 0)
		mdm.clock.SetDuration(InactivityTimer, 0)
		logger.Debug().Msg("Leaving Datamode")
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

func (mdm *Modem) getInfo() []string {
	routes := GetRoutes()
	info := []string{
		fmt.Sprintf("PiModem - Hayes compatable modem emulator"),
		fmt.Sprintf("Version: %s", version),
		fmt.Sprintf("Build: %v", buildDate),
		fmt.Sprintf("Hash: %s", gitCommit),
	}
	for name, intf := range routes {
		if !strings.HasPrefix(name, "eth") && !strings.HasPrefix(name, "wlan") {
			continue
		}
		info = append(info, fmt.Sprintf("Interface: %s", name))
		addrs, err := intf.Interface.Addrs()
		if err == nil {
			for _, addr := range addrs {
				info = append(info, fmt.Sprintf("IP Address: %v", addr))
			}
		} else {
			info = append(info, fmt.Sprintf("IP Address: %s", err.Error()))
		}
		for _, route := range intf.Routes {
			if route.Default {
				info = append(info, fmt.Sprintf("Gateway: %v", route.Gateway))
			}
		}
	}
	return info
}

func (mdm *Modem) ResetWithProfile(profile byte) error {
	var err error = nil
	if profile == 0 {
		err = mdm.reg.Load("profile0")
	} else if profile == 1 {
		err = mdm.reg.Load("profile1")
	} else {
		err = fmt.Errorf("Invalid profile")
	}
	return err
}

func (mdm *Modem) StoreProfile(profile byte) error {
	var err error = nil
	if profile == 0 {
		err = mdm.reg.Save("profile0")
	} else if profile == 1 {
		err = mdm.reg.Save("profile1")
	} else {
		err = fmt.Errorf("Invalid profile")
	}
	return err
}

func NewModem(dev string, baud uint32, address string) (*Modem, error) {
	var mdm Modem
	var err error
	mdm.serialChannel = make(chan []byte)
	mdm.CommandMode = true
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
	mdm.nvmem, err = nvmem.Open("at24c512")
	if err != nil {
		logger.Warn().Str("device", "at24c512").Msg("Cannot open nvmem device.")
	}
	mdm.reg = NewRegisters(mdm.nvmem)
	mdm.reg.SetCallback(RegGeneralBitmapOptions, mdm.bitmapOptionsRegSetCallback)
	mdm.reg.SetCallback(RegAutoAnswer, mdm.autoAnswerRegSetCallback)
	return &mdm, nil
}

func (mdm *Modem) handleCommandInput(input []byte) {
	for _, char := range input {
		if mdm.readRegister(RegStatusOptions)&Echo == Echo {
			mdm.serial.Write([]byte{char})
		}
		if string(mdm.ATHeader[:]) == "AT" || string(mdm.ATHeader[:]) == "at" {
			if char == mdm.readRegister(RegCarriageReturnChar) {
				mdm.LastCommand = append([]byte{}, mdm.Command...)
				logger.Info().Str("command", fmt.Sprintf("AT%s", mdm.Command)).Msg("Processing AT command")
				handler, err := mdm.Parse()
				for handler != nil {
					handler, err = handler(mdm)
				}
				mdm.sendCRLF()
				mdm.sendResponse(err)
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
				logger.Info().Str("command", fmt.Sprintf("AT%s", mdm.Command)).Msg("Processing AT command")
				handler, err := mdm.Parse()
				for handler != nil {
					handler, err = handler(mdm)
				}
				mdm.sendCRLF()
				mdm.sendResponse(err)
				mdm.Command = make([]byte, 0)
				mdm.ATHeader = [2]byte{0, 0}
			default:
				mdm.ATHeader[0] = char
			}
		}
	}

}

func (mdm *Modem) Start() {
	logger.Info().Str("version", version).Msg("Starting PiModem")
	logger.Info().Str("date", buildDate).Str("commit", gitCommit).Msg("Build")
	if mdm.nvmem != nil {
		powerup, err := mdm.nvmem.ReadCell("powerup")
		if err == nil {
			mdm.ResetWithProfile(byte(powerup.(uint16) & 0xff))
		}
	}
	go mdm.readSerial()
	for {
		select {
		case <-WatchPin("shutdown"):
			h := uint(syscall.LINUX_REBOOT_CMD_HALT)
			syscall.Sync()
			syscall.Reboot(int(h))
		case event := <-WatchPin("dtr"):
			logger.Debug().Int("value", int(event.Value)).Msg("DTR event")
			if event.Value == 0 {
				delay := time.Duration(mdm.readRegister(RegDelayToDTR)) * 10 * time.Millisecond
				mdm.clock.SetDuration(DTRTimer, delay)
			} else {
				mdm.clock.SetDuration(DTRTimer, 0)
			}
		case <-mdm.clock.GetTimer(DTRTimer):
			mdm.clock.SetDuration(DTRTimer, 0)
			mode := (mdm.readRegister(RegGeneralBitmapOptions) >> 3) & 3
			switch mode {
			case 0:
				logger.Debug().Msg("DTR - Ignored")
			case 1:
				logger.Debug().Msg("DTR - Return to command mode")
				mdm.setDataMode(false)
			case 2:
				logger.Debug().Msg("DTR - Hangup")
				mdm.line.Hangup()
				mdm.setDataMode(false)
			case 3:
				logger.Debug().Msg("DTR - Soft Reset (not implemented)")
			}
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
				if mdm.guardExpired || (mdm.breakCount > 0 && len(bytes) <= 3-mdm.breakCount) {
					for _, ch := range bytes {
						if ch != mdm.readRegister(RegEscapeSeqChar) {
							mdm.breakCount = 0
							break
						}
						mdm.breakCount++
					}
				} else {
					mdm.breakCount = 0
				}
				mdm.resetGuardTimer()
			}
			logger.Debug().Hex("data", bytes).Msgf("%d bytes recieved on serial", len(bytes))
		case bytes := <-mdm.line.Data:
			logger.Debug().Hex("data", bytes).Msgf("%d bytes recieved over tcp", len(bytes))
			mdm.resetInactivityTimer()
			mdm.serial.Write(bytes)
		case response := <-mdm.line.Response:
			mdm.sendResponse(response)
			if response.code == Connect {
				mdm.setDataMode(true)
				AssertPin("dsr")
				AssertPin("dcd")
			} else if response.code == NoCarrier {
				mdm.setDataMode(false)
				if mdm.readRegister(RegGeneralBitmapOptions)&0x40 == 0x40 {
					DeassertPin("dsr")
				}
				if mdm.readRegister(RegGeneralBitmapOptions)&0x20 == 0x20 {
					DeassertPin("dcd")
				}
			} else if response.code == Ring {
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
