package main

import (
	"fmt"
	"sync"
)

// Function to handle even numbers
func handleEven(numbers <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	for number := range numbers {
		if number%2 == 0 {
			fmt.Printf("Even number: %d\n", number)
		}
	}
}

// Function to handle odd numbers
func handleOdd(numbers <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	for number := range numbers {
		if number%2 != 0 {
			fmt.Printf("Odd number: %d\n", number)
		}
	}
}

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Channels for even and odd numbers
	evenChannel := make(chan int)
	oddChannel := make(chan int)

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	// Start goroutines to handle even and odd numbers
	wg.Add(2)
	go handleEven(evenChannel, &wg)
	go handleOdd(oddChannel, &wg)

	// Send numbers to respective channels
	for _, num := range numbers {
		if num%2 == 0 {
			evenChannel <- num
		} else {
			oddChannel <- num
		}
	}

	// Close channels after sending all numbers
	close(evenChannel)
	close(oddChannel)

	// Wait for goroutines to complete
	wg.Wait()
	fmt.Println("All goroutines have finished.")
}
