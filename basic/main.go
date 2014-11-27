package main

import (
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
	gbot := gobot.NewGobot()

	adaptor := sphero.NewSpheroAdaptor("sphero", "/dev/tty.Sphero-PRW-RN-SPP")
	driver := sphero.NewSpheroDriver(adaptor, "sphero")

	work := func() {
		driver.Stop()
		driver.SetRGB(0, 255, 255) // Blue
		time.Sleep(5 * time.Second)

		driver.SetRGB(255, 140, 0) // Orange
		time.Sleep(1 * time.Second)
		driver.Roll(100, 0)
		time.Sleep(2 * time.Second)
		driver.Stop()
		time.Sleep(2 * time.Second)

		driver.SetRGB(148, 0, 211) // Purple
		time.Sleep(1 * time.Second)
		driver.Roll(100, 180)
		time.Sleep(2 * time.Second)
		driver.Stop()
		time.Sleep(2 * time.Second)

		driver.SetRGB(0, 255, 255) // Blue

		driver.Stop()
		os.Exit(1)
	}

	robot := gobot.NewRobot(
		"sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
