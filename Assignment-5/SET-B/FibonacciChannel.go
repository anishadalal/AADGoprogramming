package main

import (
	"fmt"
)

// Function to generate Fibonacci series and send it to the channel
func generateFibonacci(n int, ch chan int) {
	first, second := 0, 1
	for i := 0; i < n; i++ {
		ch <- first
		next := first + second
		first = second
		second = next
	}
	close(ch) // Close the channel after sending all values
}

func main() {
	// Create a channel for Fibonacci numbers
	fibonacciChannel := make(chan int)

	// Number of terms to generate
	n := 10

	// Start a goroutine to generate Fibonacci series
	go generateFibonacci(n, fibonacciChannel)

	// Read and print Fibonacci series from the channel
	fmt.Println("Fibonacci Series:")
	for fib := range fibonacciChannel {
		fmt.Println(fib)
	}
}
