package main

import "fmt"

type Writer interface {
	Write()
}

type Reader interface {
	Read()
}

type File struct{}

func (f File) Write() {
	fmt.Println("Writing to file...")
}

func (f File) Read() {
	fmt.Println("Reading from file...")
}

func main() {
	var w Writer = File{}
	var r Reader = File{}

	w.Write()
	r.Read()
}
