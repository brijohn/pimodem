package main

import (
	"fmt"
	"strconv"
)

type HandlerFunc func(mdm *Modem) (HandlerFunc, error)
type CommandType byte

const (
	Normal CommandType = iota
	Extended
)

type Pair struct {
	t CommandType
	b byte
}

var commands map[Pair]HandlerFunc

func init() {
	commands = map[Pair]HandlerFunc{
		Pair{Normal, 'A'}: AnswerHandler,
		Pair{Normal, 'O'}: OnlineHandler,
		Pair{Normal, 'H'}: HangupHandler,
		Pair{Normal, 'D'}: DialHandler,
		Pair{Normal, 'E'}: EchoHandler,
		Pair{Normal, 'Q'}: QuietHandler,
		Pair{Normal, 'V'}: VerboseHandler,
		Pair{Normal, 'S'}: CurrentRegisterHandler,
		Pair{Normal, '?'}: QueryRegisterHandler,
		Pair{Normal, '='}: SetRegisterHandler,
	}
}

func (mdm *Modem) getNextInt(min byte, max byte, req bool, def byte) (byte, error) {
	var data []byte
	var err error
	var val uint64 = uint64(def)
	for len(mdm.Command) != 0 && mdm.Command[0] >= 48 && mdm.Command[0] <= 57 {
		data = append(data, mdm.Command[0])
		mdm.Command = mdm.Command[1:]
	}
	if len(data) > 0 {
		val, err = strconv.ParseUint(string(data), 10, 8)
		if err != nil || (byte(val) < min || byte(val) > max) {
			err = NewResponse(Error, "Invalid value")
		}
	} else if req {
		err = NewResponse(Error, "Missing required parameter")
	}
	return byte(val), err
}

func (mdm *Modem) Parse() (HandlerFunc, error) {
	var ct = Normal
	var cmd byte
	if len(mdm.Command) == 0 {
		return nil, NewResponse(Ok, "Command Successful")
	}
	cmd, mdm.Command = mdm.Command[0], mdm.Command[1:]
	switch cmd {
	case '&':
		ct = Extended
		if len(mdm.Command) == 0 {
			return nil, NewResponse(Error, "Parse error")
		}
		cmd, mdm.Command = mdm.Command[0], mdm.Command[1:]
		fallthrough
	default:
		handler, exists := commands[Pair{ct, cmd}]
		if exists {
			return handler, nil
		}
	}
	return nil, NewResponse(Error, "Unknown Command")
}

func EchoHandler(mdm *Modem) (HandlerFunc, error) {
	val, err := mdm.getNextInt(0, 1, false, 1)
	if err != nil {
		return nil, err
	}
	options := mdm.readRegister(RegStatusOptions)
	if val == 1 {
		mdm.reg.Write(RegStatusOptions, options|Echo)
	} else {
		mdm.reg.Write(RegStatusOptions, options&^Echo)
	}
	return mdm.Parse()
}

func QuietHandler(mdm *Modem) (HandlerFunc, error) {
	val, err := mdm.getNextInt(0, 1, false, 0)
	if err != nil {
		return nil, err
	}
	options := mdm.readRegister(RegStatusOptions)
	if val == 1 {
		mdm.reg.Write(RegStatusOptions, options|Quiet)
	} else {
		mdm.reg.Write(RegStatusOptions, options&^Quiet)
	}
	return mdm.Parse()
}

func VerboseHandler(mdm *Modem) (HandlerFunc, error) {
	val, err := mdm.getNextInt(0, 1, false, 1)
	if err != nil {
		return nil, err
	}
	options := mdm.readRegister(RegStatusOptions)
	if val == 1 {
		mdm.reg.Write(RegStatusOptions, options|Verbose)
	} else {
		mdm.reg.Write(RegStatusOptions, options&^Verbose)
	}
	return mdm.Parse()
}

func DialHandler(mdm *Modem) (HandlerFunc, error) {
	var number string
	if len(mdm.Command) > 0 {
		_, number = mdm.Command[0], string(mdm.Command[1:])
		mdm.Command = []byte{}
		mdm.line.SetRaw(mdm.readRegister(RegLineMode) != 0)
		err := mdm.line.Dial(number)
		if err != nil {
			return nil, err
		}
		mdm.setDataMode(true)
	}
	return nil, NewResponse(Connect, "Connecting to remote host")
}

func AnswerHandler(mdm *Modem) (HandlerFunc, error) {
	mdm.line.Pickup()
	return mdm.Parse()
}

func OnlineHandler(mdm *Modem) (HandlerFunc, error) {
	_, err := mdm.getNextInt(0, 1, false, 0)
	if err != nil {
		return nil, err
	}
	if mdm.line.Busy() {
		mdm.setDataMode(true)
	}
	return mdm.Parse()
}

func HangupHandler(mdm *Modem) (HandlerFunc, error) {
	_, err := mdm.getNextInt(0, 1, false, 0)
	if err != nil {
		return nil, err
	}
	mdm.line.Hangup()
	return mdm.Parse()
}

func CurrentRegisterHandler(mdm *Modem) (HandlerFunc, error) {
	sreg, err := mdm.getNextInt(0, 255, true, 0)
	if err != nil {
		return nil, err
	}
	mdm.reg.SetCurrent(SRegister(sreg))
	return mdm.Parse()
}

func QueryRegisterHandler(mdm *Modem) (HandlerFunc, error) {
	mdm.sendCRLF()
	val, err := mdm.reg.ReadCurrent()
	if err != nil {
		return nil, NewResponse(Error, err.Error())
	}
	//str := strconv.FormatUint(uint64(val), 10)
	str := fmt.Sprintf("%03d", val)
	mdm.serial.Write([]byte(str))
	return mdm.Parse()
}

func SetRegisterHandler(mdm *Modem) (HandlerFunc, error) {
	val, err := mdm.getNextInt(0, 255, true, 0)
	if err != nil {
		return nil, err
	}
	mdm.reg.WriteCurrent(byte(val))
	return mdm.Parse()
}
