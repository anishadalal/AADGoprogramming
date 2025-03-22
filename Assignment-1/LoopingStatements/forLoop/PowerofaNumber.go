package main

import "fmt"

func main() {
	var base, exponent, result int
	result = 1

	fmt.Print("Enter the base number: ")
	fmt.Scanln(&base)

	fmt.Print("Enter the exponent: ")
	fmt.Scanln(&exponent)

	for i := 1; i <= exponent; i++ {
		result *= base
	}

	fmt.Printf("%d raised to the power of %d is: %d\n", base, exponent, result)
}
