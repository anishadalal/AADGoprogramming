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

	var sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	average := float64(sum) / float64(n)
	fmt.Printf("The average of the elements in the array is: %.2f\n", average)
}
