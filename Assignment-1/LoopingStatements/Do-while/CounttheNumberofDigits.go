package main

import "fmt"

func main() {
	var num int
	var count int = 0

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	for {
		count++
		num /= 10
		if num == 0 {
			break
		}
	}

	fmt.Printf("The number has %d digits.\n", count)
}
