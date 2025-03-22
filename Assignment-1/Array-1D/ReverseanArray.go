package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scanln(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}

	// Reversing the array
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println("Reversed array:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
