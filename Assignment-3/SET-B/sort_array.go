package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{34, 12, 56, 89, 1, 45, 23} // Example array

	// Sorting the array in ascending order
	sort.Ints(arr)

	fmt.Println("Sorted Array in Ascending Order:", arr)
}
