package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter three numbers: ")
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Println("The largest number is", a)
	} else if b >= a && b >= c {
		fmt.Println("The largest number is", b)
	} else {
		fmt.Println("The largest number is", c)
	}
}
