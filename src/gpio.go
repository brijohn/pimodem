package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var gpio = GPIO{}

type PinEvent struct {
	Name  string
	Value int
	Time  time.Time
}

type Pin struct {
	Number       int
	Watch        int
	Notification chan PinEvent `yaml:"-"`
}

type GPIO struct {
	Hardware string
	Pins     map[string]Pin
}

func init() {
	err := yaml.Unmarshal(readConfigFile(), &gpio)
	if err != nil {
		panic(err)
	}
	initGPIO()
}

func getConfigPath() string {
	filename := "/proc/device-tree/hat/custom_0"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		progname := filepath.Base(os.Args[0])
		if os.Getenv("XDG_CONFIG_HOME") != "" {
			filename = filepath.Join(os.Getenv("XDG_CONFIG_HOME"), progname, "pins.yaml")
		} else {
			filename = filepath.Join(os.Getenv("HOME"), ".config", progname, "pins.yaml")
		}
	}
	return filename
}

func readConfigFile() []byte {
	contents, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		return make([]byte, 0)
	}
	return contents
}

func AssertPin(name string) {
	pin, ok := gpio.Pins[name]
	if ok {
		pin.High()
	}
}

func DeassertPin(name string) {
	pin, ok := gpio.Pins[name]
	if ok {
		pin.Low()
	}
}

func TogglePin(name string) {
	pin, ok := gpio.Pins[name]
	if ok {
		pin.Toggle()
	}
}

func PinValue(name string) byte {
	pin := gpio.Pins[name]
	return pin.Value()
}

func WatchPin(name string) <-chan PinEvent {
	pin, ok := gpio.Pins[name]
	if ok {
		if pin.Notification == nil {
			pin.Notification = make(chan PinEvent)
			gpio.Pins[name] = pin
		}
		return pin.Notification
	}
	return nil
}
