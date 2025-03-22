package main

import "fmt"

func main() {
	var str1, str2 string

	// Ask the user to enter two strings
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)

	// Check if str1 is a substring of str2
	if contains(str2, str1) {
		fmt.Println(str1, "is a substring of", str2)
	} else {
		fmt.Println(str1, "is not a substring of", str2)
	}
}

// Function to check if str1 is a substring of str2
func contains(str2, str1 string) bool {
	return len(str1) <= len(str2) && len(str1) > 0 && str2[:len(str1)] == str1
}
