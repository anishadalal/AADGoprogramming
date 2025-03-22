package main

import "fmt"

func updateArray(arr *[3]int) {
	arr[0] = 100
}

func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Println("Before update:", arr)

	updateArray(&arr) // Pass array's pointer to function
	fmt.Println("After update:", arr)
}
