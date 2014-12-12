// http://html-color-codes.info/
// Conn: bluez-test-serial -i hci0 68:86:E7:05:D5:02 - Sphero2
// Conn: bluez-test-serial -i hci0 00:06:66:4F:3D:A6 - Sphero1
package main

import (
	"time"

	"github.com/goinggo/robot/potato/ball"
)

var Stage1Color = [3]uint8{19, 240, 23}  // Green
var Stage2Color = [3]uint8{240, 240, 19} // Yello
var Stage3Color = [3]uint8{246, 11, 66}  // Red

func main() {
	robot := ball.Robot{
		Name:  "BLUE-SP2-68:86:E7:05:D5:02",
		Port:  "/dev/rfcomm0",
		Color: [3][3]uint8{Stage1Color, Stage2Color, Stage3Color},
		Blink: [3]time.Duration{1000 * time.Millisecond, 1000 * time.Millisecond, 500 * time.Millisecond},
	}

	robot.Run()
}
