// +build linux

/*

1) Original: Copyright (c) 2005-2008 Dustin Sallings <dustin@spy.net>.

2) Mods: Copyright (c) 2012 Schleibinger Geräte Teubert u. Greim GmbH
<info@schleibinger.com>. Blame: Jan Mercl

3) PiModem Mods: Copyright (C) 2018 Brian Johnson

All rights reserved.  Use of this source code is governed by a MIT-style
license that can be found in the LICENSE file.

*/

package main

import (
	"errors"
	"golang.org/x/sys/unix"
	"os"
	"unsafe"
)

var baudRates = map[uint32]uint32{
	300:    unix.B300,
	600:    unix.B600,
	1200:   unix.B1200,
	2400:   unix.B2400,
	4800:   unix.B4800,
	9600:   unix.B9600,
	19200:  unix.B19200,
	38400:  unix.B38400,
	57600:  unix.B57600,
	115200: unix.B115200,
}

type SerialPort struct {
	f    *os.File
	name string
	baud uint32
}

func OpenSerial(dev string, rate uint32) (*SerialPort, error) {
	var f *os.File
	var err error
	defer func() {
		if err != nil && f != nil {
			f.Close()
		}
	}()
	f, err = os.OpenFile(dev, unix.O_RDWR|unix.O_NOCTTY|unix.O_ASYNC, 0666)
	if err != nil {
		return nil, err
	}

	r, exist := baudRates[rate]
	if !exist {
		err = errors.New("Invalid Buad Rate")
		return nil, err
	}
	com := SerialPort{f, dev, rate}

	t := unix.Termios{
		Iflag:  unix.IGNBRK,
		Cflag:  unix.CS8 | unix.CREAD | unix.CLOCAL | unix.CRTSCTS | r,
		Cc:     [19]uint8{unix.VMIN: 1, unix.VTIME: 0},
		Ispeed: r,
		Ospeed: r,
	}
	err = com.Flush()
	if err != nil {
		return nil, err
	}
	err = com.SetAttributes(&t)
	if err != nil {
		return nil, err
	}

	return &com, nil
}

func (p *SerialPort) Read(b []byte) (int, error) {
	return p.f.Read(b)
}

func (p *SerialPort) Write(b []byte) (int, error) {
	return p.f.Write(b)
}

func (p *SerialPort) Close() error {
	return p.f.Close()
}

func (p *SerialPort) Flush() error {
	fd := p.f.Fd()
	if _, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(unix.TCFLSH),
		uintptr(unix.TCIFLUSH),
	); errno != 0 {
		return errno
	}
	return nil
}

func (p *SerialPort) SetAttributes(tio *unix.Termios) error {
	fd := p.f.Fd()
	if _, _, errno := unix.Syscall6(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(unix.TCSETS),
		uintptr(unsafe.Pointer(tio)),
		0,
		0,
		0,
	); errno != 0 {
		return errno
	}
	return nil
}

func (p *SerialPort) GetAttributes() (*unix.Termios, error) {
	var tio unix.Termios
	fd := p.f.Fd()
	if _, _, errno := unix.Syscall6(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(unix.TCGETS),
		uintptr(unsafe.Pointer(&tio)),
		0,
		0,
		0,
	); errno != 0 {
		return nil, errno
	}
	return &tio, nil
}

func (p *SerialPort) SetModemStatus(status uint32) error {
	fd := p.f.Fd()
	current, err := p.GetModemStatus()
	if err != nil {
		return err
	}
	current &^= (unix.TIOCM_RTS | unix.TIOCM_DTR)
	current |= status

	if _, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(unix.TIOCMSET),
		uintptr(unsafe.Pointer(&current)),
	); errno != 0 {
		return errno
	}
	return nil

}

func (p *SerialPort) GetModemStatus() (uint32, error) {
	var status uint32
	fd := p.f.Fd()

	if _, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(unix.TIOCMGET),
		uintptr(unsafe.Pointer(&status)),
	); errno != 0 {
		return 0, errno
	}
	return status, nil
}

func (p *SerialPort) SetFlowControl(status uint32) error {
	tio, err := p.GetAttributes()
	if err != nil {
		return err
	}
	tio.Cflag &^= (unix.CRTSCTS | unix.IXON | unix.IXOFF)
	tio.Cflag |= status
	err = p.SetAttributes(tio)
	if err != nil {
		return err
	}
	return nil
}
