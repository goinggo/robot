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
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		robot := ball.Robot{
			WG:    &wg,
			Name:  "BLUE-SP2-68:86:E7:05:D5:02",
			Port:  "/dev/rfcomm1",
			Color: [3]uint8{0, 0, 255},
			Coll:  ball.Collision{[]uint8{0x01, 0xF0, 0xF0, 0x80, 0x80, 0x60}, 0x02, 0x12},
		}
		robot.Run()
	}()

	go func() {
		robot := ball.Robot{
			WG:    &wg,
			Name:  "GREEN-SP1-00:06:66:4F:3D:A6",
			Port:  "/dev/rfcomm2",
			Color: [3]uint8{51, 102, 0},
			Coll:  ball.Collision{[]uint8{0x01, 0x40, 0x40, 0x80, 0x80, 0x60}, 0x02, 0x12},
		}
		robot.Run()
	}()

	wg.Wait()
}
