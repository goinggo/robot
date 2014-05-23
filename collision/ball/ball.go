package ball

import (
	"log"
	"sync"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
)

// Robot tracks collisions
type Robot struct {
	Outs int
}

// Run has the gaming logic
func (b *Robot) Run(waitGroup *sync.WaitGroup, name string, port string, color [3]uint8) {
	// Set up the adapter.
	sa := new(gobotSphero.SpheroAdaptor)
	sa.Name = "Sphero"
	sa.Port = port

	// New sphero driver.
	sd := gobotSphero.NewSphero(sa)
	sd.Name = "Sphero" + name
	sd.SetStabilization(true)

	// Channel to talk to the device.
	talk := make(chan string)

	// Work function is provided to gorobot to control the robot.
	work := func() {
		// Tell the robot to run the pause logic on collisions.
		gobot.On(sd.Events["Collision"], func(data interface{}) {
			b.Outs++
			log.Println(name, "Collision Detected - Pausing", b.Outs)

			if b.Outs == 3 {
				talk <- "shutdown"
			} else {
				talk <- "pause"
			}
		})

		// Shutdown the game after a minute.
		gobot.After("60s", func() {
			log.Println("GAME OVER")
			talk <- "shutdown"
		})

		// Starting up the robot.
		log.Println(name, "Starting Robot", color)
		sd.SetRGB(color[0], color[1], color[2])

	shutdown:
		for {
			select {
			case command := <-talk:
				switch command {
				// Pause command stops the robot and turns the color red.
				case "pause":
					log.Println(name, "Pausing Robot")
					sd.Stop()
					sd.SetRGB(220, 20, 60) // Red
					time.Sleep(1 * time.Second)

				// Shutdown command breaks from the work loop.
				case "shutdown":
					break shutdown
				}

			// Every second choose a random direction and roll fast.
			case <-time.After(2 * time.Second):
				direction := uint16(gobot.Rand(360))
				log.Println(name, "Chaning Direction", direction)
				sd.SetRGB(color[0], color[1], color[2])
				sd.Roll(255, direction)
			}
		}

		// On shutdown stop the robot and change the color to blue.
		log.Println(name, "Shutting Down Robot")
		sd.Stop()

		for i := 0; i < 3; i++ {
			sd.SetRGB(220, 20, 60) // Red
			time.Sleep(1 * time.Second)
			sd.SetRGB(255, 255, 255)
			time.Sleep(1 * time.Second)
		}

		sd.SetRGB(220, 20, 60) // Red
		time.Sleep(1 * time.Second)

		log.Println("DONE")
		waitGroup.Done()
	}

	// Control structure for the code and events.
	robot := gobot.Robot{
		Connections: []gobot.Connection{sa},
		Devices:     []gobot.Device{sd},
		Work:        work,
	}

	robot.Start()
}
