package main

import "fmt"

func main() {
	var num, divisor int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Print("Enter a divisor: ")
	fmt.Scan(&divisor)

	if divisor == 0 {
		fmt.Println("Divisor cannot be zero.")
	} else if num%divisor == 0 {
		fmt.Printf("The number %d is divisible by %d.\n", num, divisor)
	} else {
		fmt.Printf("The number %d is not divisible by %d.\n", num, divisor)
	}
}
