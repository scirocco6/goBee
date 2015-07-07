package main

import (
	"bufio"
	"errors"
	"net"
	"os"
	"strings"

	"bitbucket.org/scirocco6/icb"
)

// ReadFromUser is the main user input thread
func ReadFromUser(connection net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")
		if message == "" {
			continue
		}

		packet, err := parse(message)
		if err == nil && packet != nil {
			packet.SendTo(connection)
		} else if err != nil {
			PrintToScreen(err.Error())
		}
	}
}

func parse(message string) (*icb.Packet, error) {
	if strings.HasPrefix(message, "/") { // handle commands
		message = message[1:]
		command := strings.SplitN(message, " ", 3)

		switch command[0] {
		case "beep":
			{
				return beep(command)
			}
		case "m": // send a private message to a user
			{
				return privateMessage(command)
			}
		case "w": // obtain a listing of who is on
			{
				return globalWho(command), nil
			}
		case "g": // join a group
			{
				return join(command)
			}
		case "q": // quit
			{
				cleanExit()
			}
		default:
			{
				return nil, errors.New("Unrecognized command \n'/" + message + "'\n")
			}
		}
	}
	// by defualt send a public message to the channel
	return publicMessage(message), nil
}

func beep(parameters []string) (*icb.Packet, error) {
	if len(parameters) != 2 {
		return nil, errors.New("Usage: /beep nick")
	}

	return icb.CreatePacket("beep", parameters[1]), nil
}

func publicMessage(message string) *icb.Packet {

	return icb.CreatePacket("public", message)
}

func privateMessage(parameters []string) (*icb.Packet, error) {
	if len(parameters) != 3 {
		return nil, errors.New("Usage: /m nick message")
	}

	return icb.CreatePacket("private", parameters[1], parameters[2]), nil
}

func globalWho(parameters []string) *icb.Packet {
	if len(parameters) == 1 { // obtain listing of who is in every group
		return icb.CreatePacket("global_who")
	}
	// listing for who is in a particular group or group a user is in
	return icb.CreatePacket("local_who", parameters[1])
}

func join(parameters []string) (*icb.Packet, error) {
	if len(parameters) != 1 {
		return nil, errors.New("Usage: /g group")
	}

	return icb.CreatePacket("join", parameters[1]), nil
}
