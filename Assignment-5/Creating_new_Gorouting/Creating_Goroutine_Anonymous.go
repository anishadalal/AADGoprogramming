package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Inside Goroutine:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println("Inside Main function:", i)
		time.Sleep(500 * time.Millisecond)
	}

	// Adding sleep to allow the goroutine to complete execution
	time.Sleep(3 * time.Second)
}
