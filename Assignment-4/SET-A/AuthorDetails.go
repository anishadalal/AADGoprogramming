package main

import "fmt"

// Author struct with Name and Age fields
type Author struct {
	Name string
	Age  int
}

// Show method for Author struct
func (a Author) Show() {
	fmt.Printf("Author: %s, Age: %d\n", a.Name, a.Age)
}

func main() {
	// Create an instance of Author
	author := Author{Name: "Satyajit Ray", Age: 96}

	// Call the Show method on the author instance
	author.Show()
}
