package main

import (
	"fmt"
	"strconv"
)

// Function to calculate the sum of squares of digits
func calculateSquares(number string, result chan int) {
	sum := 0
	for _, digit := range number {
		num, _ := strconv.Atoi(string(digit))
		sum += num * num
	}
	result <- sum
}

// Function to calculate the sum of cubes of digits
func calculateCubes(number string, result chan int) {
	sum := 0
	for _, digit := range number {
		num, _ := strconv.Atoi(string(digit))
		sum += num * num * num
	}
	result <- sum
}

func main() {
	// Input number
	number := "123"

	// Channels for squares and cubes
	squaresChannel := make(chan int)
	cubesChannel := make(chan int)

	// Goroutines for calculation
	go calculateSquares(number, squaresChannel)
	go calculateCubes(number, cubesChannel)

	// Retrieve results from channels
	sumOfSquares := <-squaresChannel
	sumOfCubes := <-cubesChannel

	// Print results
	fmt.Println("Sum of squares =", sumOfSquares)
	fmt.Println("Sum of cubes =", sumOfCubes)
	fmt.Println("Final sum of squares and cubes =", sumOfSquares+sumOfCubes)
}
