package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 1 is running...")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 completed.")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 2 is running...")
	time.Sleep(1 * time.Second)
	fmt.Println("Task 2 completed.")
}

func main() {
	var wg1, wg2 sync.WaitGroup

	wg1.Add(1)
	go task1(&wg1)

	wg2.Add(1)
	go task2(&wg2)

	wg1.Wait() // Wait for task1
	wg2.Wait() // Wait for task2

	fmt.Println("Both tasks completed.")
}
