package main

import (
	//	"fmt"
	"net"
	"time"
)

const (
	LineOnHook = iota
	LineOffHook
	LineInUse
)

type Line struct {
	net.Conn
	net.Listener

	Data     chan []byte
	Response chan *ModemResponseError

	state    int
	paused   bool
	raw      bool
	terminal NVT
}

func NewLine(address string) (*Line, error) {
	var err error
	conn := Line{}
	conn.state = LineOnHook
	conn.raw = false
	conn.paused = true
	conn.Data = make(chan []byte)
	conn.Response = make(chan *ModemResponseError)
	conn.Listener, err = net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	go conn.accept()
	return &conn, nil
}

func (l *Line) readRemote() {
	buffer := make([]byte, 0, 256)

	for !l.OnHook() {
		if !l.paused {
			bytes := make([]byte, 256)
			read, err := l.Read(bytes)
			if err == nil {
				if l.paused {
					buffer = append(buffer, bytes[:read]...)
				} else {
					if len(buffer) > 0 {
						l.Data <- buffer
						buffer = buffer[:0]
					}
					l.Data <- bytes[:read]
				}
			} else {
				time.Sleep(time.Millisecond * 100)
				l.Response <- NewResponse(NoCarrier, err.Error())
				break
			}
		} else {
			time.Sleep(time.Millisecond * 100)
		}
	}
	l.Disconnect()
}

func (l *Line) answerCall(conn net.Conn) bool {
	if l.OffHook() {
		if !l.raw {
			l.Conn = NewNVT(conn, make(map[TelnetOption]OptionQueue))
		} else {
			l.Conn = conn
		}
		l.state = LineInUse
		go l.readRemote()
		return true
	}
	return false
}

func (l *Line) accept() {
accept:
	for {
		conn, err := l.Listener.Accept()
		if err != nil {
			break
		}
		if l.Busy() {
			conn.Close()
		} else {
			if l.answerCall(conn) {
				l.Response <- NewResponse(Connect, "Connecting to remote host")
				continue accept
			}
			for r := 0; r < 15; r++ {
				playAudio("ring.wav", nil)
				l.Response <- NewResponse(Ring, "Phone ringing; Incoming call")
				d := 0
				for d < 4000 {
					time.Sleep(20 * time.Millisecond)
					d += 20
					if l.answerCall(conn) {
						l.Response <- NewResponse(Connect, "Connecting to remote host")
						continue accept
					}
				}
			}
			conn.Close()
		}
	}
}

func (l *Line) Dial(address string) error {
	if l.Busy() || l.OffHook() {
		return NewResponse(Busy, "Line Busy")
	}
	playAudio("dial.wav", nil)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return NewResponse(NoCarrier, err.Error())
	}
	if !l.raw {
		l.Conn = NewNVT(conn, make(map[TelnetOption]OptionQueue))
	} else {
		l.Conn = conn
	}
	l.state = LineInUse
	go l.readRemote()
	return nil
}

func (l *Line) Hangup() {
	l.Disconnect()
}

func (l *Line) Pickup() {
	if l.OnHook() {
		l.state = LineOffHook
	}
}

func (l *Line) Pause(paused bool) {
	l.paused = paused
}

func (l *Line) SetRaw(raw bool) {
	l.raw = raw
}

func (l *Line) Read(b []byte) (int, error) {
	return l.Conn.Read(b)

}

func (l *Line) Write(b []byte) (int, error) {
	return l.Conn.Write(b)
}

func (l *Line) Disconnect() {
	l.state = LineOnHook
	if l.Conn != nil {
		l.Conn.Close()
	}
}

func (l *Line) OnHook() bool {
	return l.state == LineOnHook
}

func (l *Line) OffHook() bool {
	return l.state == LineOffHook
}

func (l *Line) Busy() bool {
	return l.state == LineInUse
}
