package main

import "fmt"

func main() {
	var ch string
	fmt.Print("Enter a character: ")
	fmt.Scan(&ch)

	if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {
		fmt.Println("The character is a vowel.")
	} else {
		fmt.Println("The character is a consonant.")
	}
}
