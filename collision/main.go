// Scan: hcitool scan
// Pair: bluez-simple-agent hci0 ADDRESS
// Conn: bluez-test-serial -i hci0 68:86:E7:05:D5:02 - YELLOW (NEW)
// Conn: bluez-test-serial -i hci0 00:06:66:4F:3D:A6 - BLUE (OLD)
package main

import (
	"log"
	"os"
	"sync"

	"github.com/goinggo/robot/collision/ball"
	"github.com/hybridgroup/gobot/platforms/sphero"
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
			CC: sphero.CollisionConfig{
				Method: 0x01,
				Xt:     0x80,
				Yt:     0x80,
				Xs:     0x80,
				Ys:     0x80,
				Dead:   0x60,
			},
		}
		robot.Run()
	}()

	go func() {
		robot := ball.Robot{
			WG:    &wg,
			Name:  "GREEN-SP1-00:06:66:4F:3D:A6",
			Port:  "/dev/rfcomm2",
			Color: [3]uint8{51, 102, 0},
			CC: sphero.CollisionConfig{
				Method: 0x01,
				Xt:     0x40,
				Yt:     0x40,
				Xs:     0x80,
				Ys:     0x80,
				Dead:   0x60,
			},
		}
		robot.Run()
	}()

	wg.Wait()
}
