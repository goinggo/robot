package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"

	"os"
	"time"
)

func main() {

	spheroAdaptor := new(gobotSphero.SpheroAdaptor)
	spheroAdaptor.Name = "Sphero"
	spheroAdaptor.Port = "/dev/tty.Sphero-Ball1-RN-SPP"

	sphero := gobotSphero.NewSphero(spheroAdaptor)
	sphero.Name = "Sphero"

	work := func() {
		sphero.Stop()
		sphero.SetRGB(0, 255, 255) // Blue
		time.Sleep(5 * time.Second)

		sphero.SetRGB(255, 140, 0) // Orange
		time.Sleep(1 * time.Second)
		sphero.Roll(100, 0)
		time.Sleep(2 * time.Second)
		sphero.Stop()
		time.Sleep(2 * time.Second)

		sphero.SetRGB(148, 0, 211) // Purple
		time.Sleep(1 * time.Second)
		sphero.Roll(100, 180)
		time.Sleep(2 * time.Second)
		sphero.Stop()
		time.Sleep(2 * time.Second)

		sphero.SetRGB(0, 255, 255) // Blue

		sphero.Stop()
		os.Exit(1)
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{spheroAdaptor},
		Devices:     []gobot.Device{sphero},
		Work:        work,
	}

	robot.Start()
}
