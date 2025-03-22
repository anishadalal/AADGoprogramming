package main

import "fmt"

func main() {
	var n, target int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scanln(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}

	fmt.Print("Enter the element to search for: ")
	fmt.Scanln(&target)

	found := false
	for i := 0; i < n; i++ {
		if arr[i] == target {
			found = true
			fmt.Printf("Element %d found at index %d.\n", target, i)
			break
		}
	}

	if !found {
		fmt.Println("Element not found in the array.")
	}
}
