package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scanln(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}

	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	fmt.Printf("The sum of the elements in the array is: %d\n", sum)
}
