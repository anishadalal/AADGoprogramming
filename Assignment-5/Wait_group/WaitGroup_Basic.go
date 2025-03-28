package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark the goroutine as done
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done.\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add to the WaitGroup counter
		go worker(i, &wg)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers completed.")
}
