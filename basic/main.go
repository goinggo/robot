// Scan: hcitool scan
// Pair: bluez-simple-agent hci0 ADDRESS
// Conn: bluez-test-serial -i hci0 68:86:E7:05:D5:02 - YELLOW (NEW)
// gort bluetooth connect 68:86:E7:05:D5:02
// Conn: bluez-test-serial -i hci0 00:06:66:4F:3D:A6 - BLUE (OLD)
// gort bluetooth connect 00:06:66:4F:3D:A6
package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
	g := gobot.NewGobot()
	a := sphero.NewSpheroAdaptor("sphero", "/dev/rfcomm0")
	d := sphero.NewSpheroDriver(a, "sphero")

	work := func() {
		d.Stop()
		d.SetRGB(0, 255, 255) // Blue
		time.Sleep(5 * time.Second)

		d.SetRGB(255, 140, 0) // Orange
		time.Sleep(1 * time.Second)
		d.Roll(255, 0)
		time.Sleep(5 * time.Second)
		d.Stop()
		time.Sleep(2 * time.Second)

		d.SetRGB(148, 0, 211) // Purple
		time.Sleep(1 * time.Second)
		d.Roll(225, 180)
		time.Sleep(5 * time.Second)
		d.Stop()
		time.Sleep(2 * time.Second)

		d.SetRGB(255, 255, 255) // Blue
		time.Sleep(1 * time.Second)

		d.Stop()
	}

	robot := gobot.NewRobot(
		"sphero",
		[]gobot.Connection{a},
		[]gobot.Device{d},
		work,
	)

	g.AddRobot(robot)
	g.Start()
}
