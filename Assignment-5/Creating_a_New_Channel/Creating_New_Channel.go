package main

import (
	"fmt"
)

func sendMessage(ch chan string) {
	ch <- "Hello from Goroutine!" // Sending data into the channel
}

func main() {
	// Creating a new channel of type string
	messageChannel := make(chan string)

	// Starting a goroutine to send a message
	go sendMessage(messageChannel)

	// Receiving data from the channel
	message := <-messageChannel
	fmt.Println("Received:", message)
}
