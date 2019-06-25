package main

import (
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

	state   int
	ringing bool
	Volume  int
	Speaker bool

	paused   bool
	raw      bool
	terminal NVT
}

func NewLine(address string) (*Line, error) {
	var err error
	conn := Line{}
	conn.state = LineOnHook
	conn.ringing = false
	conn.raw = false
	conn.Volume = 2
	conn.Speaker = true
	conn.paused = true
	conn.Data = make(chan []byte)
	conn.Response = make(chan *ModemResponseError, 1)
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
				if l.Established() {
					l.Response <- NewResponse(NoCarrier, err.Error())
				}
				break
			}
		} else {
			time.Sleep(time.Millisecond * 100)
		}
	}
	l.Disconnect()
}

func (l *Line) volumeLevel() int {
	if !l.Speaker || l.Volume == 0 {
		return 0
	}
	return l.Volume
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
		l.ringing = false
		conn, err := l.Listener.Accept()
		if err != nil {
			break
		}
		if l.Busy() {
			conn.Close()
		} else {
			l.ringing = true
			if l.answerCall(conn) {
				l.Response <- NewResponse(Connect, "Connecting to remote host")
				continue accept
			}
			for r := 0; r < 15; r++ {
				playAudio("ring.wav", l.volumeLevel(), nil)
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

func (l *Line) Dial(address string, timeout byte) {
	if l.Busy() {
		l.Response <- NewResponse(Busy, "Line Busy")
	} else {
		playAudio("v34-33600bps.wav", l.volumeLevel(), nil)
		conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
		if err != nil {
			l.Response <- NewResponse(NoAnswer, err.Error())
		} else {
			if !l.raw {
				l.Conn = NewNVT(conn, make(map[TelnetOption]OptionQueue))
			} else {
				l.Conn = conn
			}
			l.state = LineInUse
			l.Response <- NewResponse(Connect, "Connecting to remote host")
			go l.readRemote()
		}
	}
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

func (l *Line) Ringing() bool {
	return l.ringing
}

func (l *Line) Established() bool {
	return l.state == LineInUse
}

func (l *Line) Busy() bool {
	return l.Ringing() || l.Established()
}
