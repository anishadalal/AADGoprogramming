package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter the value of N: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("The sum of the first %d natural numbers is: %d\n", n, sum)
}
