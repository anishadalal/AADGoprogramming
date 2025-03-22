package main

import "fmt"

type Number struct {
	value int
}

func (n Number) valueReceiver() {
	n.value += 10
	fmt.Println("Inside Value Receiver:", n.value)
}

func (n *Number) pointerReceiver() {
	n.value += 10
	fmt.Println("Inside Pointer Receiver:", n.value)
}

func main() {
	n := Number{value: 20}

	n.valueReceiver()
	fmt.Println("After Value Receiver:", n.value)

	n.pointerReceiver()
	fmt.Println("After Pointer Receiver:", n.value)
}
