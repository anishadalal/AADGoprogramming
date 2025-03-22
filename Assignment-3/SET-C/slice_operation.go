package main

import "fmt"

func main() {
	// Initialize a slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial Slice:", slice)

	// Append elements
	slice = append(slice, 6, 7)
	fmt.Println("After Appending 6 and 7:", slice)

	// Remove an element (e.g., remove the element at index 2)
	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println("After Removing Element at Index 2:", slice)

	// Copy a slice
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	fmt.Println("Copied Slice:", copySlice)

	// Length and Capacity
	fmt.Println("Length of Slice:", len(slice))
	fmt.Println("Capacity of Slice:", cap(slice))
}
