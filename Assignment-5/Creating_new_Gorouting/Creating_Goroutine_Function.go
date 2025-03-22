package main

import (
	"fmt"
	"time"
)

func printMessage() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Goroutine")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage() // Running function as a Goroutine

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Main function")
		time.Sleep(500 * time.Millisecond)
	}

	// Adding sleep to allow goroutine to finish execution
	time.Sleep(3 * time.Second)
}
