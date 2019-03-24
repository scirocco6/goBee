package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var screenMutex = &sync.Mutex{} // global lock for screen output
var sane *terminal.State

func main() {
	catchSignals()
	defer cleanExit()

	sane, _ = terminal.GetState(0)
	defer terminal.Restore(0, sane)

	fmt.Println("Connecting...")
	ConnectToServer()

	runtime.GOMAXPROCS(runtime.NumCPU())
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

// This should be reworked as a gofunc reading from a channel
// PrintToScreen locks the screen mutex then prints the string it is passed
func PrintToScreen(message string) {
	screenMutex.Lock()
	inSane, _ := terminal.GetState(0)
	terminal.Restore(0, sane)
	fmt.Println(message)
	terminal.Restore(0, inSane)
	screenMutex.Unlock()
}

func cleanExit() {
	os.Exit(1)
}
