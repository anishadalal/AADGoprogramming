package main

import "fmt"

func increment(val *int) {
	*val += 1
	fmt.Printf("Inside increment function: val = %d\n", *val)
}

func main() {
	num := 5
	fmt.Printf("Before increment: num = %d\n", num)
	increment(&num)
	fmt.Printf("After increment: num = %d\n", num)
}
