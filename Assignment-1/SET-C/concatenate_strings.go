package main

import "fmt"

func concatenateStrings(str1, str2 *string) string {
	// Concatenate the strings pointed by str1 and str2
	return *str1 + *str2
}

func main() {
	// Declare two strings
	str1 := "Hello, "
	str2 := "World!"

	// Call concatenateStrings function with pointers to str1 and str2
	result := concatenateStrings(&str1, &str2)

	// Print the concatenated result
	fmt.Println("Concatenated String:", result)
}
