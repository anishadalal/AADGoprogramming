package main

import "fmt"

func main() {
	var num, sum int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	for num > 0 {
		sum += num % 10
		num /= 10
	}
	fmt.Printf("The sum of the digits is: %d\n", sum)
}
