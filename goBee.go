package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/rthornton128/goncurses"
)

//var term *goncurses.Window // the global hook to the terminal
var screenMutex = &sync.Mutex{} // global lock for screen output

func main() {
	catchSignals()
	defer cleanExit()

	fmt.Println("Connecting...")
	connection := ConnectToServer()

	//result := ReadPacketFromConnection(connection)
	//message := result.Decode()
	//
	//	if strings.HasPrefix(message, "[=Login=]") {
	//		PrintToScreen(message)
	//	} else {
	//		PrintToScreen(message)
	//		PrintToScreen("Unable to login to server")
	//		cleanExit()
	//	}

	//	term = initializeCurses()
	//	term.Clear()
	//	fmt.Println("normal print")
	//	term.Println("curses print")
	//	term.Refresh()
	//_ = term.GetChar()
	//goncurses.UnGetChar(c) buzz GetChar is in window but ungetchar is in gc??? and they are different types???

	runtime.GOMAXPROCS(2)
	ReadFromServer(connection)
	ReadFromUser(connection)
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

	//goncurses.CBreak(true)
	return term
}

// PrintToScreen locks the screen mutex then prints the string it is passed
func PrintToScreen(message string) {
	screenMutex.Lock()
	//term.Println(message)
	//term.Refresh()

	fmt.Println(message)
	screenMutex.Unlock()
}

func cleanExit() {
	// restore terminal on program exit
	goncurses.CBreak(false)
	goncurses.Echo(true)
	goncurses.End()
	os.Exit(1)
}
