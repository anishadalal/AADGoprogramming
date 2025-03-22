package main

import "fmt"

func main() {
	var n, num, largest int
	fmt.Print("How many numbers? ")
	fmt.Scanln(&n)

	i := 1
	for i <= n {
		fmt.Printf("Enter number %d: ", i)
		fmt.Scanln(&num)

		if num > largest {
			largest = num
		}
		i++
	}
	fmt.Printf("The largest number is: %d\n", largest)
}
