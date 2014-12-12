package ball

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

// Robot tracks collisions.
type Robot struct {
	Name  string
	Port  string
	Color [3][3]uint8
	Blink [3]time.Duration
	Stage int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot) Run() {
	gbot := gobot.NewGobot()
	sa := sphero.NewSpheroAdaptor(r.Name, r.Port)
	sd := sphero.NewSpheroDriver(sa, r.Name)
	sd.SetBackLED(0)

	reset := func() {
		log.Println("Resetting Game")
		sd.Stop()
		sd.SetRGB(255, 255, 255)
		r.Stage = 0

		for i := 5; i >= 1; i-- {
			fmt.Println("Game Starting In", i)
			time.Sleep(1 * time.Second)
		}

		fmt.Println("STAGE : Stage[1]")
	}

	over := func() {
		fmt.Println("Game Over")
		sd.SetRGB(220, 20, 60)
	}

	play := func() {
		reset()
		var gc int

		for {
			time.Sleep(r.Blink[r.Stage])

			if r.Stage == 0 {
				sd.SetRGB(r.Color[r.Stage][0], r.Color[r.Stage][1], r.Color[r.Stage][2])
			} else {
				sd.SetRGB(255, 255, 255)
				time.Sleep(r.Blink[r.Stage])
				sd.SetRGB(r.Color[r.Stage][0], r.Color[r.Stage][1], r.Color[r.Stage][2])
			}

			gc++
			if gc <= 3 {
				continue
			}

			if rand.Intn(10) == 3 || gc == 10 {
				gc = 0

				r.Stage++
				fmt.Printf("STAGE : Stage[%d]\n", r.Stage+1)

				if r.Stage == 3 {
					over()

					reader := bufio.NewReader(os.Stdin)
					fmt.Println("Hit <ENTER> To Play Again")
					v, _ := reader.ReadString('\n')
					if v == "q\n" {
						fmt.Println("Thanks For Playing")
						break
					}

					reset()
				}
			}
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
