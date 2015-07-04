package main

import (
	"net"

	"bitbucket.org/scirocco6/icb"
)

// ReadFromServer is the main server read loop
func ReadFromServer(connection net.Conn) {
	go func() {
		for {
			var packet icb.Packet

			length := ReadLengthByteFromConnection(connection)
			buffer := ReadLengthBytesFromServer(length, connection)
			//		fmt.Printf("length: %v\n  data:\n\t%v\n\t\"%s\"\n", length, buffer, buffer)

			packet.Write(buffer)

			message := packet.Decode()
			//term.Println(message)
			//term.Refresh()
			PrintToScreen(message)
		}
	}()
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
