package main

import "fmt"

func main() {
	var number int

	// Ask the user to enter a number
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	// Check if the number is a single digit
	if number >= -9 && number <= 9 {
		fmt.Println("The number is a single digit.")
	} else {
		fmt.Println("The number is not a single digit.")
	}
}
