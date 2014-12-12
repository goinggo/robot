package ball

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

var red uint8
var green uint8
var blue uint8

// Robot tracks collisions.
type Robot struct {
	Name string
	Port string
}

func (r *Robot) Run() {
	gbot := gobot.NewGobot()
	sa := sphero.NewSpheroAdaptor(r.Name, r.Port)
	sd := sphero.NewSpheroDriver(sa, r.Name)
	sd.SetBackLED(0)

	play := func() {
		for {
			fmt.Println("Enter color numbers")
			fmt.Scanf("%d %d %d", &red, &green, &blue)
			fmt.Println(red, green, blue)

			if red == 0 && green == 0 && blue == 0 {
				fmt.Println("Safe To Shutdown")
				break
			}

			sd.SetRGB(red, green, blue)
		}
	}

	robot := gobot.NewRobot(
		"sphero",
		[]gobot.Connection{sa},
		[]gobot.Device{sd},
		play,
	)

	gbot.AddRobot(robot)
	gbot.Start()

	fmt.Println("Shutting Down Robot")
}
