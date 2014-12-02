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
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		blue := [3]uint8{0, 0, 255}
		var robot ball.Robot
		robot.Run(&waitGroup, "sphero", "/dev/rfcomm0", blue)
	}()

	//go func() {
	//	green := [3]uint8{51, 102, 0}
	//	var robot ball.Robot
	//	robot.Run(&waitGroup, "Erick", "/dev/tty.Sphero-Erick-RN-SPP", green)
	//}()

	waitGroup.Wait()
}
