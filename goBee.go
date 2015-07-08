package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

var screenMutex = &sync.Mutex{} // global lock for screen output

func main() {
	catchSignals()
	defer cleanExit()

	fmt.Println("Connecting...")
	ConnectToServer()

	runtime.GOMAXPROCS(2)
	ReadFromServer()
	ReadFromUser()
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

// PrintToScreen locks the screen mutex then prints the string it is passed
func PrintToScreen(message string) {
	screenMutex.Lock()
	fmt.Println(message)
	screenMutex.Unlock()
}

func cleanExit() {
	// restore terminal on program exit

	os.Exit(1)
}
