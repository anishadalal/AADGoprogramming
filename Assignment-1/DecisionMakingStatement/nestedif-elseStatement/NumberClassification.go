package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num > 0 {
		if num%2 == 0 {
			fmt.Println("The number is positive and even.")
		} else {
			fmt.Println("The number is positive and odd.")
		}
	} else if num < 0 {
		if num%2 == 0 {
			fmt.Println("The number is negative and even.")
		} else {
			fmt.Println("The number is negative and odd.")
		}
	} else {
		fmt.Println("The number is zero.")
	}
}
