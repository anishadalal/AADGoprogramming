package main

import (
	"fmt"
)

func main() {
	arr := []int{34, 12, 56, 89, 1, 45, 23} // Example array
	largest, smallest := arr[0], arr[0]

	for _, num := range arr {
		if num > largest {
			largest = num
		}
		if num < smallest {
			smallest = num
		}
	}

	fmt.Println("Array:", arr)
	fmt.Println("Largest Number:", largest)
	fmt.Println("Smallest Number:", smallest)
}
