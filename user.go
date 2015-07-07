package main

import (
	"bufio"
	"fmt"
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
					if len(command) == 2 {
						packet = icb.CreatePacket("beep", command[1])
					} else {
						PrintToScreen("Usage: /beep nick")
					}
				}
			case "m": // send a private message to a user
				{
					if len(command) == 3 {
						packet = icb.CreatePacket("private", command[1], command[2])
					} else {
						PrintToScreen("Usage: /m nick message")
					}
				}
			case "w": // obtain a listing of who is on
				{
					if len(command) == 1 { // obtain listing of who is in every group
						packet = icb.CreatePacket("global_who")
					} else { // listing for who is in a particular group or group a user is in
						packet = icb.CreatePacket("local_who", command[1])
					}
				}
			case "g": // join a group
				{
					if len(command) == 1 {
						packet = icb.CreatePacket("join", command[1])
					} else {
						PrintToScreen("Usage: /g group")
					}
				}
			case "q":
				{
					cleanExit()
				}
			default:
				{
					fmt.Printf("Unrecognized command \n'/%s'\n", message)
					continue
				}
			}
		} else { // by defualt send a public message to the channel
			packet = icb.CreatePacket("public", message)
		}
		packet.SendTo(connection)
	}
}
