package main

import "fmt"

// Recursive function to calculate sum of digits
func sumOfDigits(num int) int {
	if num == 0 {
		return 0
	}
	return num%10 + sumOfDigits(num/10)
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Please enter a positive number.")
	} else {
		result := sumOfDigits(number)
		fmt.Printf("The sum of digits of %d is %d\n", number, result)
	}
}
