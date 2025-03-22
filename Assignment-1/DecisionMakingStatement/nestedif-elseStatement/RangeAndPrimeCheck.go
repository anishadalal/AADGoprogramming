package main

import "fmt"

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Check if the number is within the range of 1 to 100
	if number >= 1 && number <= 100 {
		// Nested if-else to check if the number is prime
		if isPrime(number) {
			fmt.Println("The number is within the range and is prime.")
		} else {
			fmt.Println("The number is within the range and is not prime.")
		}
	} else {
		fmt.Println("The number is outside the range of 1 to 100.")
	}
}
