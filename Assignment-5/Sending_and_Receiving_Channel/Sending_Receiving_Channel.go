package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Hello, Channel!" // Sending data to the channel
}

func main() {
	// Creating a channel of type string
	messageChannel := make(chan string)

	// Starting a goroutine to send data
	go sendData(messageChannel)

	// Receiving data from the channel
	receivedMessage := <-messageChannel
	fmt.Println("Received:", receivedMessage)
}
