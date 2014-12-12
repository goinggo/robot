// http://html-color-codes.info/
// Conn: bluez-test-serial -i hci0 68:86:E7:05:D5:02 - Sphero2
// Conn: bluez-test-serial -i hci0 00:06:66:4F:3D:A6 - Sphero1
package main

import "github.com/goinggo/robot/color/ball"

func main() {
	robot := ball.Robot{
		Name: "BLUE-SP2-68:86:E7:05:D5:02",
		Port: "/dev/rfcomm0",
	}

	robot.Run()
}
