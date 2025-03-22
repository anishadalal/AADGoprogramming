package main

import "fmt"

func printSlice(s []int) {
	fmt.Println("Slice:", s)
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	printSlice(slice)
}
