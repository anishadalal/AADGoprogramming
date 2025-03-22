package main

import "fmt"

func rectangle(length, breadth int) (area int, perimeter int) {
	area = length * breadth
	perimeter = 2 * (length + breadth)
	return
}

func main() {
	a, p := rectangle(5, 3)
	fmt.Println("Area:", a, "Perimeter:", p)
}
