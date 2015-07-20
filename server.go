package main

import "bitbucket.org/scirocco6/icb"

// ConnectToServer returns a connection to the icb server
func ConnectToServer() {
	icb.Connect("default.icb.net", "7326")
	_ = ReadLengthByteFromConnection // read the protocol version then ignore it :)

	loginPacket := icb.CreatePacket("login", "goBee", "goBee7", "goGroup", "login", "\000")
	loginPacket.Send()
}

// ReadFromServer is the main server read loop
func ReadFromServer() {
	go func() {
		for {
			packet := ReadPacketFromConnection()
			message := packet.Decode()

			PrintToScreen(message)
		}
	}()
}

// ReadPacketFromConnection will read a complete packet from the connection and return it
func ReadPacketFromConnection() icb.Packet {
	var packet icb.Packet

	length := ReadLengthByteFromConnection()
	buffer := ReadLengthBytesFromServer(length)
	//		fmt.Printf("length: %v\n  data:\n\t%v\n\t\"%s\"\n", length, buffer, buffer)

	packet.Write(buffer)
	return packet
}

// ReadLengthByteFromConnection reads the next byte from the connection and returns it as an int
func ReadLengthByteFromConnection() int {
	var length = make([]byte, 1)
	icb.Connection.Read(length)

	return int(length[0])
}

// ReadLengthBytesFromServer reads an entire packet of length from the connection returns a byte array
func ReadLengthBytesFromServer(length int) []byte {
	var buffer = make([]byte, length)
	var bytesRead = 0

	for bytesRead < length {
		r, _ := icb.Connection.Read(buffer[bytesRead:])
		bytesRead += r
	}

	return buffer
}
