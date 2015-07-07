package main

import (
	"log"
	"net"
	"time"

	"bitbucket.org/scirocco6/icb"
)

// ConnectToServer returns a connection to the icb server
func ConnectToServer() net.Conn {
	connection, err := icb.Connect("default.icb.net", "7326")
	if err == nil {
		connection.SetDeadline(time.Time{}) //   do not time out on i/o operations

		_ = ReadLengthByteFromConnection // read the protocol version then ignore it :)

		loginPacket := icb.CreatePacket("login", "goBee", "goBee6", "goGroup", "login", "\000")
		loginPacket.SendTo(connection)

		return connection
	}

	log.Fatal("Connecting:", err)
	cleanExit()
	return nil
}

// ReadFromServer is the main server read loop
func ReadFromServer(connection net.Conn) {
	go func() {
		for {
			packet := ReadPacketFromConnection(connection)
			message := packet.Decode()

			PrintToScreen(message)
		}
	}()
}

// ReadPacketFromConnection will read a complete packet from the connection and return it
func ReadPacketFromConnection(connection net.Conn) icb.Packet {
	var packet icb.Packet

	length := ReadLengthByteFromConnection(connection)
	buffer := ReadLengthBytesFromServer(length, connection)
	//		fmt.Printf("length: %v\n  data:\n\t%v\n\t\"%s\"\n", length, buffer, buffer)

	packet.Write(buffer)
	return packet
}

// ReadLengthByteFromConnection reads the next byte from the connection and returns it as an int
func ReadLengthByteFromConnection(connection net.Conn) int {
	var length = make([]byte, 1)
	connection.Read(length)

	return int(length[0])
}

// ReadLengthBytesFromServer reads an entire packet of length from the connection returns a byte array
func ReadLengthBytesFromServer(length int, connection net.Conn) []byte {
	var buffer = make([]byte, length)
	var bytesRead = 0

	for bytesRead < length {
		r, _ := connection.Read(buffer[bytesRead:])
		bytesRead += r
	}

	return buffer
}
