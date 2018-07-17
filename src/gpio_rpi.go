// +build rpi

package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

func initGPIO() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	for _, pin := range gpio.Pins {
		rpio.Pin(pin.Number).Detect(rpio.Edge(pin.Watch))
	}
	go pollPins()
}

func (p *Pin) High() {
	rpio.Pin(p.Number).High()
}

func (p *Pin) Low() {
	rpio.Pin(p.Number).Low()
}

func (p *Pin) Toggle() {
	rpio.Pin(p.Number).Toggle()
}

func (p *Pin) Value() byte {
	return byte(rpio.Pin(p.Number).Read())
}

func pollPins() {
	for {
		for name, pin := range gpio.Pins {
			if rpio.Edge(pin.Watch) != rpio.NoEdge && pin.Notification != nil {
				if rpio.Pin(pin.Number).EdgeDetected() {
					pin.Notification <- PinEvent{name, int(rpio.Pin(pin.Number).Read()), time.Now()}
				}
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
}
