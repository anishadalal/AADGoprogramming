package main

import "fmt"

func main() {
	var number int

	// Ask the user to enter a number
	fmt.Print("Enter a number to print its multiplication table: ")
	fmt.Scanln(&number)

	// Print the multiplication table
	fmt.Println("Multiplication table of", number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
