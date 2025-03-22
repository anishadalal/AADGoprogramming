package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the value of N: ")
	fmt.Scanln(&n)

	fmt.Println("Even numbers from 1 to", n, ":")
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
