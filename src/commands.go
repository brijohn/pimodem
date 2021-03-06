/** commands.go
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
		Pair{Normal, 'A'}:   AnswerHandler,
		Pair{Normal, 'O'}:   OnlineHandler,
		Pair{Normal, 'H'}:   HangupHandler,
		Pair{Normal, 'D'}:   DialHandler,
		Pair{Normal, 'E'}:   EchoHandler,
		Pair{Normal, 'Q'}:   QuietHandler,
		Pair{Normal, 'V'}:   VerboseHandler,
		Pair{Normal, 'I'}:   InfoHandler,
		Pair{Normal, 'M'}:   SpeakerHandler,
		Pair{Normal, 'L'}:   VolumeHandler,
		Pair{Normal, 'X'}:   ExtendedResponseHandler,
		Pair{Normal, 'Z'}:   ResetHandler,
		Pair{Normal, 'S'}:   CurrentRegisterHandler,
		Pair{Normal, '?'}:   QueryRegisterHandler,
		Pair{Normal, '='}:   SetRegisterHandler,
		Pair{Extended, 'C'}: DCDModeHandler,
		Pair{Extended, 'D'}: DTRModeHandler,
		Pair{Extended, 'S'}: DSROverrideHandler,
		Pair{Extended, 'W'}: StoreProfileHandler,
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

func InfoHandler(mdm *Modem) (HandlerFunc, error) {
	info := mdm.getInfo()
	val, err := mdm.getNextInt(0, byte(len(info)-1), false, 255)
	if err != nil {
		return nil, err
	}
	if val == 255 {
		for _, line := range info {
			mdm.sendCRLF()
			mdm.serial.Write([]byte(line))
		}
	} else if int(val) < len(info) {
		mdm.sendCRLF()
		mdm.serial.Write([]byte(info[val]))
	}
	return mdm.Parse()
}

func DialHandler(mdm *Modem) (HandlerFunc, error) {
	var number string
	if len(mdm.Command) > 0 {
		_, number = mdm.Command[0], string(mdm.Command[1:])
		mdm.Command = []byte{}
		mdm.line.SetRaw(mdm.readRegister(RegLineMode) != 0)
		mdm.line.Dial(number, mdm.readRegister(RegWaitForCarrierDelay))
	}
	return nil, nil
}

func AnswerHandler(mdm *Modem) (HandlerFunc, error) {
	if mdm.line.Ringing() {
		mdm.line.Pickup()
	}
	return nil, NewResponse(Ok, "Command Successful")
}

func OnlineHandler(mdm *Modem) (HandlerFunc, error) {
	_, err := mdm.getNextInt(0, 1, false, 0)
	if err != nil {
		return nil, err
	}
	if mdm.line.Established() {
		mdm.setDataMode(true)
	}
	return mdm.Parse()
}

func HangupHandler(mdm *Modem) (HandlerFunc, error) {
	mode, err := mdm.getNextInt(0, 1, false, 0)
	if err != nil {
		return nil, err
	}
	switch mode {
	case 0:
		mdm.line.Hangup()
	case 1:
		mdm.line.Pickup()
	}
	return mdm.Parse()
}

func VolumeHandler(mdm *Modem) (HandlerFunc, error) {
	volume, err := mdm.getNextInt(0, 3, true, 0)
	if err != nil {
		return nil, err
	}
	mdm.setVolume(int(volume))
	return mdm.Parse()
}

func SpeakerHandler(mdm *Modem) (HandlerFunc, error) {
	speaker, err := mdm.getNextInt(0, 3, true, 0)
	if err != nil {
		return nil, err
	}
	mdm.setSpeakers(int(speaker))
	return mdm.Parse()
}

func ExtendedResponseHandler(mdm *Modem) (HandlerFunc, error) {
	_, err := mdm.getNextInt(0, 4, true, 0)
	if err != nil {
		return nil, err
	}
	return mdm.Parse()
}

func ResetHandler(mdm *Modem) (HandlerFunc, error) {
	profile, err := mdm.getNextInt(0, 1, false, 0)
	if err != nil {
		return nil, err
	}
	err = mdm.ResetWithProfile(profile)
	if err != nil {
		return nil, NewResponse(Error, err.Error())
	}
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

func DCDModeHandler(mdm *Modem) (HandlerFunc, error) {
	mode, err := mdm.getNextInt(0, 1, true, 0)
	if err != nil {
		return nil, err
	}
	val := mdm.readRegister(RegGeneralBitmapOptions) & 0xDF
	val |= (mode << 5)
	mdm.writeRegister(RegGeneralBitmapOptions, val)
	return mdm.Parse()
}

func DTRModeHandler(mdm *Modem) (HandlerFunc, error) {
	mode, err := mdm.getNextInt(0, 3, true, 0)
	if err != nil {
		return nil, err
	}
	val := mdm.readRegister(RegGeneralBitmapOptions) & 0xE7
	val |= (mode << 3)
	mdm.writeRegister(RegGeneralBitmapOptions, val)
	return mdm.Parse()
}

func DSROverrideHandler(mdm *Modem) (HandlerFunc, error) {
	mode, err := mdm.getNextInt(0, 1, true, 0)
	if err != nil {
		return nil, err
	}
	val := mdm.readRegister(RegGeneralBitmapOptions) & 0xBF
	val |= (mode << 6)
	mdm.writeRegister(RegGeneralBitmapOptions, val)
	return mdm.Parse()
}

func StoreProfileHandler(mdm *Modem) (HandlerFunc, error) {
	profile, err := mdm.getNextInt(0, 1, true, 0)
	if err != nil {
		return nil, err
	}
	err = mdm.StoreProfile(profile)
	if err != nil {
		return nil, NewResponse(Error, err.Error())
	}
	return mdm.Parse()
}
