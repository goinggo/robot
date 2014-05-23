package main

import (
	"log"
	"os"
	"sync"

	"github.com/goinggo/robot/collision/ball"
)

func inti() {
	log.SetOutput(os.Stdout)
}

func main() {
	blue := [3]uint8{0, 0, 255}
	green := [3]uint8{51, 102, 0}

	var waitGroup sync.WaitGroup

	go func() {
		var robot ball.Robot
		robot.Run(&waitGroup, "Bill", "/dev/tty.Sphero-Bill-RN-SPP", blue)
	}()

	go func() {
		var robot ball.Robot
		robot.Run(&waitGroup, "Erick", "/dev/tty.Sphero-Erick-RN-SPP", green)
	}()

	waitGroup.Wait()
}
