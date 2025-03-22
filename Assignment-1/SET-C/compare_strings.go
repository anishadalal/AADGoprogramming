package main

import "fmt"

func main() {
	var str1, str2 string

	// Ask the user to enter two strings
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)

	// Compare the two strings
	if str1 == str2 {
		fmt.Println("The strings are equal.")
	} else {
		fmt.Println("The strings are not equal.")
	}
}
