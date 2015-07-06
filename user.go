package main

import (
	"bufio"
	"net"
	"os"
	"regexp"
	"strings"

	"bitbucket.org/scirocco6/icb"
)

// ReadFromUser is the main user input thread
func ReadFromUser(connection net.Conn) {
	beep, _ := regexp.Compile("/beep[\t\n\f\r ]([^\t\n\f\r ]+)")

	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")
		if message == "" {
			PrintToScreen("spinning")
		}

		var packet icb.Packet
		if strings.HasPrefix(message, "/") {
			if beep.MatchString(message) {

			}
			// handle commands
		} else { // by defualt send a public message to the channel
			packet = icb.CreatePacket("public", message)
		}
		packet.SendTo(connection)

		PrintToScreen(message)
	}
}
