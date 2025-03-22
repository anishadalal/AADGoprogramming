package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 4) // Creating a buffered channel with capacity 4

	fmt.Println("Channel capacity:", cap(ch)) // Should print 4
}
