package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function simulating work done by a worker in the workshop
func worker(id int, wg *sync.WaitGroup, checkpoint *sync.Mutex) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	// Simulate work with random duration
	workDuration := time.Duration(id*100) * time.Millisecond
	fmt.Printf("Worker %d is working for %v...\n", id, workDuration)
	time.Sleep(workDuration)

	// Wait at the checkpoint (synchronize with others)
	checkpoint.Lock() // Locking the checkpoint
	fmt.Printf("Worker %d is at the checkpoint and ready to proceed.\n", id)
	checkpoint.Unlock() // Unlocking for the next worker

	// Simulate task completion
	fmt.Printf("Worker %d has completed its task.\n", id)
}

func main() {
	// Number of workers
	numWorkers := 5
	var wg sync.WaitGroup
	checkpoint := &sync.Mutex{} // Mutex to synchronize at the checkpoint

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each worker
		go worker(i, &wg, checkpoint)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers have completed their tasks.")
}
