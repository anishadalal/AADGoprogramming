package main

import "fmt"

// Method to copy elements of one array to another
func copyArray(src, dst []int) {
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}
}

func main() {
	// Original array
	source := []int{1, 2, 3, 4, 5}

	// Destination array (same length as source)
	destination := make([]int, len(source))

	// Call the method to copy elements from source to destination
	copyArray(source, destination)

	// Display the copied array
	fmt.Println("Source Array:", source)
	fmt.Println("Destination Array:", destination)
}
