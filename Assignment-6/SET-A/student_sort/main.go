package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

func main() {
	students := []Student{
		{"John", 85},
		{"Jane", 90},
		{"Doe", 75},
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Marks > students[j].Marks
	})

	for _, s := range students {
		fmt.Println(s.Name, s.Marks)
	}
}
