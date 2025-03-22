package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Program started...")

	// Using time.Sleep() to pause execution for 2 seconds
	time.Sleep(2 * time.Second)

	fmt.Println("2 seconds later...")

	// Another delay of 3 seconds
	time.Sleep(3 * time.Second)

	fmt.Println("Total 5 seconds delay completed.")
}
