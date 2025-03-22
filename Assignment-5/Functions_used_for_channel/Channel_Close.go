package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3) // Creating a buffered channel

	// Sending data to the channel
	ch <- 10
	ch <- 20
	ch <- 30

	close(ch) // Closing the channel

	// Receiving data from the closed channel
	for val := range ch {
		fmt.Println("Received:", val)
	}

	fmt.Println("Channel closed, no more data can be sent.")
}
