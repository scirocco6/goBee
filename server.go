package main

import (
    "fmt"
    "net"
    "bitbucket.org/scirocco6/icb"
)

func ReadFromServer(connection net.Conn) {
    for {
        var length []byte = make([]byte, 1)
        var packet icb.IcbPacket

        connection.Read(length)
        left := int(length[0])
        for left > 0 {
            buffer := ReadLengthBytesFromServer(left, connection)
            left -= len(buffer)

            packet.Write(buffer)
        }
        message := packet.Decode()
        fmt.Println(message)
    }
}

func ReadLengthBytesFromServer(length int, connection net.Conn)([]byte) {
    var buffer []byte = make([]byte, length)
    
    connection.Read(buffer)
    
    return buffer
}