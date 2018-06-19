package main

import (
	"fmt"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Port   int    `cli:"p,port" usage:"port to listen on" dft:"6400"`
	IP     string `cli:"i,ip" usage:"ip address to listen on" dft:"0.0.0.0"`
	Device string `cli:"*d,device" usage:"serial device (e.g. /dev/ttyS0)."`
	Baud   uint32 `cli:"s,speed" usage:"baud rate for serial device" dft:"9600"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		mdm, errno := NewModem(argv.Device, argv.Baud, fmt.Sprintf("%s:%d", argv.IP, argv.Port))
		if errno != nil {
			panic(errno)
		} else {
			mdm.Start()
		}
		return nil
	})
}
