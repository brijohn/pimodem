/** pimodem.go
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
	"github.com/mkideal/cli"
	"github.com/rs/zerolog"
	"io"
	"os"
)

var version string
var buildDate string
var gitCommit string

type LogLevel zerolog.Level
type Logger zerolog.Logger

func (l *Logger) Decode(s string) error {
	var output io.Writer
	switch s {
	case "console":
		output = zerolog.ConsoleWriter{Out: os.Stderr}
	case "syslog":
		output = NewSyslogWriter()
	case "journald":
		output = NewJournalWriter()
	}
	if output != nil {
		*l = Logger(zerolog.New(output).With().Timestamp().Logger())
		return nil
	}
	return fmt.Errorf("Invalid logging destination: %s", s)
}

func (l *LogLevel) Decode(s string) error {
	level, err := zerolog.ParseLevel(s)
	*l = LogLevel(level)
	return err
}

type argT struct {
	cli.Helper
	Port      int      `cli:"p,port" usage:"port to listen on" dft:"6400"`
	IP        string   `cli:"i,ip" usage:"ip address to listen on" dft:"0.0.0.0"`
	Device    string   `cli:"*d,device" usage:"serial device (e.g. /dev/ttyS0)."`
	Baud      uint32   `cli:"s,speed" usage:"baud rate for serial device" dft:"9600"`
	Level     LogLevel `cli:"level" usage:"set logging level" dft:"info"`
	LogOutput Logger   `cli:"logger" usage:"set logging destination" dft:"console"`
}

var logger zerolog.Logger

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		zerolog.SetGlobalLevel(zerolog.Level(argv.Level))
		logger = zerolog.Logger(argv.LogOutput)
		mdm, errno := NewModem(argv.Device, argv.Baud, fmt.Sprintf("%s:%d", argv.IP, argv.Port))
		if errno != nil {
			panic(errno)
		} else {
			mdm.Start()
		}
		return nil
	})
}
