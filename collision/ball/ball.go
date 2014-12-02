package ball

import (
	"log"
	"sync"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

// Robot tracks collisions.
type Robot struct {
	WG    *sync.WaitGroup
	Name  string
	Port  string
	Color [3]uint8
	CC    sphero.CollisionConfig
	Outs  int
}

// Run has the gaming logic.
func (r *Robot) Run() {
	gbot := gobot.NewGobot()

	// Set up the adapter.
	sa := sphero.NewSpheroAdaptor(r.Name, r.Port)

	// New sphero driver.
	sd := sphero.NewSpheroDriver(sa, r.Name)
	sd.SetStabilization(true)
	sd.ConfigureCollisionDetection(r.CC)

	// Channel to talk to the device.
	talk := make(chan string)

	// Work function is provided to gorobot to control the robot.
	work := func() {
		// Tell the robot to run the pause logic on collisions.
		gobot.On(sd.Event("collision"), func(data interface{}) {
			r.Outs++
			log.Println(r.Name, "Collision Detected - Pausing", r.Outs)

			if r.Outs == 3 {
				talk <- "shutdown"
			} else {
				talk <- "pause"
			}
		})

		// Shutdown the game after a minute.
		gobot.After(300*time.Second, func() {
			log.Println("GAME OVER")
			talk <- "shutdown"
		})

		// Starting up the robot.
		log.Println(r.Name, "Starting Robot")
		sd.SetRGB(r.Color[0], r.Color[1], r.Color[2])

	shutdown:
		for {
			select {
			case command := <-talk:
				switch command {
				// Pause command stops the robot and turns the color red.
				case "pause":
					//log.Println(name, "Pausing Robot")
					sd.Stop()
					sd.SetRGB(220, 20, 60) // Red
					time.Sleep(5 * time.Second)

				// Shutdown command breaks from the work loop.
				case "shutdown":
					break shutdown
				}

			// Every second choose a random direction and roll fast.
			case <-time.After(500 * time.Millisecond):
				direction := uint16(gobot.Rand(360))
				//log.Println(r.Name, "Chaning Direction", direction)
				sd.SetRGB(r.Color[0], r.Color[1], r.Color[2])
				sd.Roll(150, direction)
			}
		}

		// On shutdown stop the robot and change the color to blue.
		log.Println(r.Name, "Shutting Down Robot")
		sd.Stop()

		for i := 0; i < 3; i++ {
			sd.SetRGB(220, 20, 60) // Red
			time.Sleep(2 * time.Second)
			sd.SetRGB(r.Color[0], r.Color[1], r.Color[2])
			time.Sleep(1 * time.Second)
		}

		log.Println(r.Name, "DONE Robot")
		r.WG.Done()
	}

	robot := gobot.NewRobot(
		"sphero",
		[]gobot.Connection{sa},
		[]gobot.Device{sd},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
