package main

import (
	"fmt"
	"sync"
	"time"
)

func process(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Process %d started...\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Process %d finished.\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All processes completed.")
}
