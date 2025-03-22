package main

import (
	"fmt"
)

func main() {
	// Creating a buffered channel with capacity 2
	messageChannel := make(chan string, 2)

	// Sending messages (does not block because of buffer)
	messageChannel <- "Hello"
	messageChannel <- "World"

	fmt.Println(<-messageChannel) // Receiving first message
	fmt.Println(<-messageChannel) // Receiving second message
}
