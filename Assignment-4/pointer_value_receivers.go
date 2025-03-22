package main

import "fmt"

type Sample struct {
	value int
}

func (s Sample) show() {
	fmt.Println("Value:", s.value)
}

func (s *Sample) modify() {
	s.value = 100
}

func main() {
	s := Sample{value: 10}
	s.show()
	s.modify()
	s.show()
}
