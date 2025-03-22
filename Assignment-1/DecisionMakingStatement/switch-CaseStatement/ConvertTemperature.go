package main

import (
	"fmt"
)

func main() {
	var choice int
	var temp float64

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Convert Celsius to Fahrenheit")
		fmt.Println("2. Convert Fahrenheit to Celsius")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 3 {
			fmt.Println("Exiting the program.")
			break
		}

		fmt.Print("Enter temperature value: ")
		fmt.Scanln(&temp)

		switch choice {
		case 1:
			fahrenheit := (temp * 9 / 5) + 32
			fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", temp, fahrenheit)
		case 2:
			celsius := (temp - 32) * 5 / 9
			fmt.Printf("%.2f Fahrenheit = %.2f Celsius\n", temp, celsius)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
