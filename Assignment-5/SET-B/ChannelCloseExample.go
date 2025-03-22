package main

import "fmt"

// Function to send values to the channel
func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all data
}

func main() {
	// Create a channel
	channel := make(chan int)

	// Start a goroutine to send data to the channel
	go sendData(channel)

	// Use the for-range loop to read values from the channel
	// The loop will stop when the channel is closed and all data has been received
	fmt.Println("Reading from the channel:")
	for value := range channel {
		fmt.Println(value)
	}

	fmt.Println("Channel is closed and all data has been received.")
}
