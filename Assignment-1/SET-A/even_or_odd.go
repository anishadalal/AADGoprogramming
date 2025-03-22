package main

import "fmt"

func main() {
	var number int

	// Ask the user to enter a number
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	// Check if the number is even or odd
	if number%2 == 0 {
		fmt.Println(number, "is Even")
	} else {
		fmt.Println(number, "is Odd")
	}
}
