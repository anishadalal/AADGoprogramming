package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the value of N: ")
	fmt.Scanln(&n)

	i := 1
	for i <= n {
		fmt.Println(i)
		i++
	}
}
