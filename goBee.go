package main

import (
	//"bufio"
	//"io"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.org/scirocco6/icb"
	"github.com/rthornton128/goncurses"
)

func main() {
	//    catchSignals()
	//defer cleanExit()

	fmt.Println("Connecting...")
	connection := connectToServer()
	fmt.Println("Connected.")

	//term := initializeCurses()
	//goncurses.CBreak(true)

	ReadFromServer(connection)
	//_ = term.GetChar() // block until the user starts typing
	//goncurses.UnGetChar(c) buzz GetChar is in window but ungetchar is in gc??? and they are different types???
}

func catchSignals() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		_ = <-signalChannel
		fmt.Println("Exiting")
		cleanExit()
	}()
}

func initializeCurses() *goncurses.Window {
	term, err := goncurses.Init() // initialize go curses

	if err != nil {
		log.Fatal("init:", err)
		cleanExit()
	}

	return term
}

func connectToServer() net.Conn {
	connection, err := icb.Connect("default.icb.net", "7326")
	if err == nil {
		connection.SetDeadline(time.Time{}) //   do not time out on i/o operations

		var b = make([]byte, 1)
		connection.Read(b) // read the protocol version then ignore it :)

		loginPacket := icb.CreatePacket("login", "goBee", "goBee6", "goGroup", "login", "\000")
		loginPacket.SendTo(connection)

		// TODO: need to check results form the login attempt rather than just assuming it worked
		// on the plus side, not suceeding doesn't prevent one from chatting

		beepPacket := icb.CreatePacket("beep", "0110")
		beepPacket.SendTo(connection)

		publicPacket := icb.CreatePacket("public", "hi everyone")
		publicPacket.SendTo(connection)

		privatePacket := icb.CreatePacket("private", "0110", "hi six")
		privatePacket.SendTo(connection)

		whoPacket := icb.CreatePacket("global_who")
		whoPacket.SendTo(connection)

		return connection
	}

	log.Fatal("Connecting:", err)
	cleanExit()
	return nil
}

func cleanExit() {
	// restore terminal on program exit
	goncurses.CBreak(false)
	goncurses.Echo(true)
	goncurses.End()
	fmt.Println("Cleaning up...")
	os.Exit(1)
}
