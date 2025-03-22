package main

import (
	"fmt"
)

func main() {
	var choice int

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Display Day of the Week")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 2 {
			fmt.Println("Exiting the program.")
			break
		}

		var day int
		fmt.Print("Enter the day number (1-7): ")
		fmt.Scanln(&day)

		switch day {
		case 1:
			fmt.Println("Sunday")
		case 2:
			fmt.Println("Monday")
		case 3:
			fmt.Println("Tuesday")
		case 4:
			fmt.Println("Wednesday")
		case 5:
			fmt.Println("Thursday")
		case 6:
			fmt.Println("Friday")
		case 7:
			fmt.Println("Saturday")
		default:
			fmt.Println("Invalid day number. Please enter a number between 1 and 7.")
		}
	}
}
