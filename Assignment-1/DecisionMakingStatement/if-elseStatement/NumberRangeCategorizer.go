package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number >= -10 && number <= 10 {
		fmt.Println("The number is small.")
	} else if number > 10 && number <= 100 {
		fmt.Println("The number is medium.")
	} else {
		fmt.Println("The number is large.")
	}
}
