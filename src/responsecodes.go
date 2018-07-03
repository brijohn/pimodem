package main

import (
	"fmt"
)

type ModemResponse int

const (
	Ok ModemResponse = iota
	Ring
	NoCarrier
	Error
	NoDialtone
	Busy
	NoAnswer
	Connect
	Connect600
	Connect1200
	Connect2400
	Connect4800
	Connect9600
	Connect12000
	Connect14400
	Connect19200
	Connect38400
	Connect57600
	Connect115200
)

var modemResponseStrings = map[ModemResponse]string{
	Ok:            "OK",
	Ring:          "RING",
	NoCarrier:     "NO CARRIER",
	Error:         "ERROR",
	NoDialtone:    "NO DIALTONE",
	Busy:          "BUSY",
	NoAnswer:      "NO ANSWER",
	Connect:       "CONNECT",
	Connect600:    "CONNECT 0600",
	Connect1200:   "CONNECT 1200",
	Connect2400:   "CONNECT 2400",
	Connect4800:   "CONNECT 4800",
	Connect9600:   "CONNECT 9600",
	Connect12000:  "CONNECT 12000",
	Connect14400:  "CONNECT 14400",
	Connect19200:  "CONNECT 19200",
	Connect38400:  "CONNECT 38400",
	Connect57600:  "CONNECT 57600",
	Connect115200: "CONNECT 115200",
}

type ModemResponseError struct {
	code ModemResponse
	msg  string
}

func NewConnectResponseFromSpeed(speed int) error {
	switch speed {
	case 600:
		return NewResponse(Connect600, "Connecting to remote host @ %d bps", speed)
	case 1200:
		return NewResponse(Connect1200, "Connecting to remote host @ %d bps", speed)
	case 2400:
		return NewResponse(Connect2400, "Connecting to remote host @ %d bps", speed)
	case 4800:
		return NewResponse(Connect4800, "Connecting to remote host @ %d bps", speed)
	case 9600:
		return NewResponse(Connect9600, "Connecting to remote host @ %d bps", speed)
	case 12000:
		return NewResponse(Connect12000, "Connecting to remote host @ %d bps", speed)
	case 14400:
		return NewResponse(Connect14400, "Connecting to remote host @ %d bps", speed)
	case 19200:
		return NewResponse(Connect19200, "Connecting to remote host @ %d bps", speed)
	case 38400:
		return NewResponse(Connect38400, "Connecting to remote host @ %d bps", speed)
	case 57600:
		return NewResponse(Connect57600, "Connecting to remote host @ %d bps", speed)
	case 115200:
		return NewResponse(Connect115200, "Connecting to remote host @ %d bps", speed)
	default:
		return NewResponse(Connect, "Connecting to remote host")
	}
}

func NewResponse(code ModemResponse, msg string, param ...interface{}) *ModemResponseError {
	return &ModemResponseError{code, fmt.Sprintf(msg, param...)}
}

func (e *ModemResponseError) Error() string {
	return e.msg
}

func (e *ModemResponseError) Response() string {
	return modemResponseStrings[e.code]
}
