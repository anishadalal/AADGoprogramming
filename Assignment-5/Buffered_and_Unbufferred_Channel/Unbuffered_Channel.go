package main

import (
	"fmt"
)

func sendMessage(ch chan string) {
	fmt.Println("Sending message...")
	ch <- "Hello from Goroutine!" // Blocks until received
	fmt.Println("Message sent!")
}

func main() {
	// Creating an unbuffered channel
	messageChannel := make(chan string)

	go sendMessage(messageChannel) // Start goroutine

	// Receiving message (unblocks sendMessage)
	message := <-messageChannel
	fmt.Println("Received:", message)
}

