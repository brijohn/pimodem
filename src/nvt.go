/** nvt.go
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
	"net"
	"time"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type NVTState int

const (
	StateData NVTState = iota
	StateIAC
	StateNegotiate
	StateSB
	StateSBData
	StateSBIAC
)

type TelnetOption int

const (
	OptionBinary TelnetOption = iota
	OptionEcho
	OptionRCP
	OptionSGA
	OptionNAMS
	OptionStatus
	OptionTM
	OptionRCTE
	OptionNAOL
	OptionNAOP
	OptionNAOCRD
	OptionNAOHTS
	OptionNAOHTD
	OptionNAOFFD
	OptionNAOVTS
	OptionNAOVTD
	OptionNAOLFD
	OptionXASCII
	OptionLogout
	OptionBM
	OptionDET
	OptionSUPDUP
	OptionSUPDUPOUTPUT
	OptionSNDLOC
	OptionTTYPE
	OptionEOR
	OptionTUID
	OptionOUTMRK
	OptionTTYLOC
	Option3270REGIME
	OptionX3PAD
	OptionNAWS
	OptionTSPEED
	OptionLFLOW
	OptionLINEMODE
	OptionXDISPLOC
	OptionENVIRON
	OptionAUTHENTICATION
	OptionENCRYPT
	OptionNEWENVIRON
)

const (
	EOF byte = (236 + iota)
	SUSP
	ABORT
	EOR
	SE
	NOP
	DM
	BREAK
	IP
	AO
	AYT
	EC
	EL
	GA
	SB
	WILL
	WONT
	DO
	DONT
	IAC
)

type OptionQueue struct {
	us    byte
	him   byte
	state uint16
}

type NVT struct {
	conn    net.Conn
	buffer  []byte
	state   NVTState
	telopts map[TelnetOption]OptionQueue
	action  byte
	sb_opt  byte
	sb_data []byte
}

func NewNVT(c net.Conn, opts map[TelnetOption]OptionQueue) *NVT {
	var nvt NVT
	nvt.conn = c
	nvt.telopts = opts
	return &nvt
}

func (nvt *NVT) Read(b []byte) (int, error) {
	buf := make([]byte, len(b))
	n, err := nvt.conn.Read(buf)
	if err != nil {
		return 0, err
	}
	nvt.processBytes(buf[:n])
	n = len(nvt.buffer)
	copy(b, nvt.buffer)
	nvt.buffer = nil
	return n, nil
}

func (nvt *NVT) Write(b []byte) (int, error) {
	//escape IAC Bytes
	return nvt.conn.Write(b)
}

func (nvt *NVT) Close() error {
	return nvt.conn.Close()
}

func (nvt *NVT) LocalAddr() net.Addr {
	return nvt.conn.LocalAddr()
}

func (nvt *NVT) RemoteAddr() net.Addr {
	return nvt.conn.RemoteAddr()
}

func (nvt *NVT) SetDeadline(dl time.Time) error {
	return nvt.conn.SetDeadline(dl)
}

func (nvt *NVT) SetReadDeadline(dl time.Time) error {
	return nvt.conn.SetReadDeadline(dl)
}

func (nvt *NVT) SetWriteDeadline(dl time.Time) error {
	return nvt.conn.SetWriteDeadline(dl)
}

func (nvt *NVT) negotiate(cmd byte) {

}

func (nvt *NVT) subnegotiate() {

}

func (nvt *NVT) processBytes(bytes []byte) {
	for _, b := range bytes {
		switch nvt.state {
		case StateData:
			if b == IAC {
				nvt.state = StateIAC
			} else {
				nvt.buffer = append(nvt.buffer, b)
			}
		case StateIAC:
			switch b {
			case IAC:
				nvt.buffer = append(nvt.buffer, b)
				nvt.state = StateData
			case WILL, WONT, DO, DONT: //WILL/WONT/DO/DONT
				nvt.action = b
				nvt.state = StateNegotiate
			case SB: //SB
				nvt.state = StateSB
			default:
				//send IAC Event
				fmt.Printf("IAC Command: %X\n", b)
				nvt.state = StateData
			}
		case StateNegotiate:
			//do negotiation
			nvt.negotiate(b)
			fmt.Printf("Negotiate: %X %X\n", nvt.action, b)
			nvt.state = StateData
		case StateSB:
			nvt.sb_opt = b
			nvt.state = StateSBData
		case StateSBData:
			if b == IAC {
				nvt.state = StateSBIAC
			} else {
				nvt.sb_data = append(nvt.sb_data, b)

			}
		case StateSBIAC:
			switch b {
			case IAC:
				nvt.sb_data = append(nvt.sb_data, b)
				nvt.state = StateSBData
			case SE:
				fmt.Printf("Subnegotiate: %X %x\n", nvt.sb_opt, nvt.sb_data)
				nvt.subnegotiate()
				//handle subnegotiate
				nvt.state = StateData
				nvt.sb_data = nil
			default:
				fmt.Printf("Unexpected byte after IAC inside SB: %d\n", b)
				nvt.subnegotiate()
				//handle subnegotiate
				nvt.state = StateData
				nvt.sb_data = nil
			}
		}
	}
}
