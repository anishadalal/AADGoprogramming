package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

func main() {
	double := multiplier(2)
	fmt.Println("Double of 5:", double(5))
}
