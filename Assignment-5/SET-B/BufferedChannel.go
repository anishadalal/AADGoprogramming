package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 5
	bufferedChannel := make(chan int, 5)

	// Store values in the channel
	bufferedChannel <- 10
	bufferedChannel <- 20
	bufferedChannel <- 30
	bufferedChannel <- 40
	bufferedChannel <- 50

	// Print the capacity and length of the channel
	fmt.Printf("Initial Channel Capacity: %d\n", cap(bufferedChannel))
	fmt.Printf("Initial Channel Length: %d\n", len(bufferedChannel))

	// Read values from the channel
	fmt.Println("Reading values from the channel:")
	for i := 0; i < 5; i++ {
		value := <-bufferedChannel
		fmt.Printf("Value: %d\n", value)
	}

	// Print the modified length of the channel
	fmt.Printf("Modified Channel Length: %d\n", len(bufferedChannel))
}
