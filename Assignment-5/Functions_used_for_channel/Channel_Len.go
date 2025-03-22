package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5) // Creating a buffered channel with capacity 5

	// Sending data
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("Current length of channel:", len(ch)) // Should print 3
}
