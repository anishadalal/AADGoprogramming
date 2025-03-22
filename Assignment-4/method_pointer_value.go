package main

import "fmt"

type Example struct {
	count int
}

// Method with a value receiver
func (e Example) show() {
	fmt.Println("Count:", e.count)
}

// Method with a pointer receiver
func (e *Example) increase() {
	e.count++
}

func main() {
	e := Example{count: 10}

	e.show()     // Calls method with value receiver
	e.increase() // Calls method with pointer receiver
	e.show()     // Should reflect the change
}
