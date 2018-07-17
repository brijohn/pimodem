// +build !rpi

package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

const NumGPIO = 28

var pins []byte
var oldPins []byte

func initGPIO() {
	file, err := os.OpenFile("/dev/shm/registers", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	file.Truncate(NumGPIO)
	defer file.Close()
	pins, err = syscall.Mmap(int(file.Fd()), 0, NumGPIO, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	oldPins = make([]byte, NumGPIO)
	copy(oldPins, pins)
	go pollPins()
}

func (p *Pin) High() {
	pins[p.Number] = 1
}

func (p *Pin) Low() {
	pins[p.Number] = 0
}

func (p *Pin) Toggle() {
	if pins[p.Number] == 1 {
		p.Low()
	} else {
		p.High()
	}
}

func (p *Pin) Value() byte {
	return pins[p.Number]
}

func pollPins() {
	for {
		for name, pin := range gpio.Pins {
			if pin.Watch != 0 && pin.Notification != nil {
				if oldPins[pin.Number] != pins[pin.Number] {
					fmt.Printf("Detected Edge Change on %s - (%d -> %d)\n", name, oldPins[pin.Number], pins[pin.Number])
					if (pin.Watch&1) == 1 && pins[pin.Number] == 1 {
						pin.Notification <- PinEvent{name, int(pins[pin.Number]), time.Now()}
					}
					if (pin.Watch&2) == 2 && pins[pin.Number] == 0 {
						pin.Notification <- PinEvent{name, int(pins[pin.Number]), time.Now()}
					}
				}
			}
			oldPins[pin.Number] = pins[pin.Number]
		}
		time.Sleep(20 * time.Millisecond)
	}
}
