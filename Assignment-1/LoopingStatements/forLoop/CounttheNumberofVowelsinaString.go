package main

import "fmt"

func main() {
	var input string
	var count int

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	for i := 0; i < len(input); i++ {
		switch string(input[i]) {
		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
			count++
		}
	}

	fmt.Printf("The number of vowels in the string is: %d\n", count)
}
