package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
	gbot := gobot.NewGobot()

	mqttAdaptor := mqtt.NewMqttAdaptor("server", "tcp://test.mosquitto.org:1883", "ball")

	adaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-PRW-AMP-SPP")
	ball := sphero.NewSpheroDriver(adaptor, "sphero")

	var move struct {
		deg   uint16
		speed uint8
	}

	work := func() {
		ball.SetBackLED(255)

		mqttAdaptor.On("/sphero/deg", func(data []byte) {
			deg, _ := strconv.Atoi(string(data))
			fmt.Println("deg:", deg)
			move.deg = uint16(deg)
		})
		mqttAdaptor.On("/sphero/speed", func(data []byte) {
			speed, _ := strconv.Atoi(string(data))
			fmt.Println("speed:", speed)
			move.speed = uint8(speed)
		})

		gobot.Every(100*time.Millisecond, func() {
			ball.Roll(move.speed, move.deg)
		})
	}

	robot := gobot.NewRobot("mqttBot",
		[]gobot.Connection{adaptor, mqttAdaptor},
		[]gobot.Device{ball},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
