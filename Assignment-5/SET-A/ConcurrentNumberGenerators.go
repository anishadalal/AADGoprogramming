package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to generate numbers with a random delay
func generateNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	for i := 0; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond) // Random delay between 0 and 250 ms
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	// Launch 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go generateNumbers(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All goroutines have finished.")
}
