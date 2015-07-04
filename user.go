package main

import (
	"bufio"
	"os"
)

// ReadFromUser is the main user input thread
func ReadFromUser() {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		PrintToScreen(message)
	}
}
