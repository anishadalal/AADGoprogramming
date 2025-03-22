package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new file
	file, err := os.Create("hello_world.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	// Write "Hello World" to the file
	_, err = file.WriteString("Hello World")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and written successfully.")
}
