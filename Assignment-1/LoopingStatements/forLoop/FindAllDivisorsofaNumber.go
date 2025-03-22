package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number to find its divisors: ")
	fmt.Scanln(&num)

	fmt.Printf("Divisors of %d are: ", num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
