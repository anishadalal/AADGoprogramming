package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
	fmt.Printf("Inside swap function: a = %d, b = %d\n", *a, *b)
}

func main() {
	x, y := 10, 20
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}
