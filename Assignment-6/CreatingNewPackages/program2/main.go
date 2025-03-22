package main

import (
	"CreatingNewPackages/program2/package2"
	"fmt"
)

func main() {
	sum := package2.Add(10, 5)
	fmt.Println("Sum:", sum)
}
