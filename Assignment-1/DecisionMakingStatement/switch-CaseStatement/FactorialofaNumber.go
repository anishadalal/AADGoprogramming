package main

import (
	"fmt"
)

func main() {
	var choice int

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Find the Factorial of a Number")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 2 {
			fmt.Println("Exiting the program.")
			break
		}

		var num int
		fmt.Print("Enter a number: ")
		fmt.Scanln(&num)

		if num < 0 {
			fmt.Println("Factorial is not defined for negative numbers.")
			continue
		}

		factorial := 1
		switch {
		case num == 0:
			factorial = 1
		default:
			for i := 1; i <= num; i++ {
				factorial *= i
			}
		}
		fmt.Printf("The factorial of %d is %d\n", num, factorial)
	}
}
