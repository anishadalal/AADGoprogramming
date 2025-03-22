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

	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Printf("The maximum element in the array is: %d\n", max)
}
