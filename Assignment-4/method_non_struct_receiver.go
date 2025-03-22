package main

import "fmt"

type MyInt int

func (m MyInt) square() int {
	return int(m * m)
}

func main() {
	var num MyInt = 5
	fmt.Println("Square:", num.square())
}
