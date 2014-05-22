package main

import (
	"log"
	"os"

	"github.com/robot/collision/ball"
)

func inti() {
	log.SetOutput(os.Stdout)
}

func main() {
	blue := [3]uint8{0, 0, 255}
	green := [3]uint8{51, 102, 0}

	if len(os.Args) == 2 {
		var robot ball.Robot

		color := green
		if os.Args[1] == "Erick" {
			color = blue
		}

		robot.Run(os.Args[1], "/dev/tty.Sphero-"+os.Args[1]+"-RN-SPP", color)
	}

	log.Println(len(os.Args))
}
