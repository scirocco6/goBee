package main

import (
	"bufio"
	"net"
	"os"
	"strings"

	"bitbucket.org/scirocco6/icb"
)

// ReadFromUser is the main user input thread
func ReadFromUser(connection net.Conn) {
	//beep, _ := regexp.Compile("/beep[\t\n\f\r ]([^\t\n\f\r ]+)")
	//command := regexp.MustCompile("/([^\t\n\f\r ]+)[\t\n\f\r ]*(.*)")

	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")
		if message == "" {
			continue
		}

		var packet icb.Packet
		if strings.HasPrefix(message, "/") { // handle commands
			message = message[1:]
			command := strings.SplitN(message, " ", 3)

			switch command[0] {
			case "beep":
				{
					if len(command) > 1 { // allow /beep user message even though no message is sent
						packet = icb.CreatePacket("beep", command[1])
					}
				}
			case "m": // send a private message to a user
				{
					if len(command) == 3 {
						packet = icb.CreatePacket("private", command[1], command[2])
					}
				}
			default:
				{
					continue
				}
			}
		} else { // by defualt send a public message to the channel
			packet = icb.CreatePacket("public", message)
		}
		packet.SendTo(connection)

		PrintToScreen(message)
	}
}
