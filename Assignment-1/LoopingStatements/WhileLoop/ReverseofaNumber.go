package main

import "fmt"

func main() {
	var num, reverse int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	for num > 0 {
		reverse = reverse*10 + num%10
		num /= 10
	}
	fmt.Printf("The reverse of the number is: %d\n", reverse)
}
