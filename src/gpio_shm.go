// +build !rpi

/** gpio_shm.go
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
