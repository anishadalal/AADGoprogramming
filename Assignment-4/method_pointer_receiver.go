package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) increment() {
	c.value++
}

func main() {
	c := Counter{value: 10}
	c.increment()
	fmt.Println("Counter Value:", c.value)
}
