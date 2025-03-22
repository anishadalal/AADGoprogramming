package main

import "fmt"

type Author struct {
	name  string
	books int
}

// Method with a receiver of Author type
func (a Author) details() {
	fmt.Println("Author:", a.name, "Books Published:", a.books)
}

func main() {
	a := Author{name: "Mark", books: 5}
	a.details()
}
