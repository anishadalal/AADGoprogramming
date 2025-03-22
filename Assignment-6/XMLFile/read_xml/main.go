package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Student struct {
	Name  string `xml:"name"`
	Marks int    `xml:"marks"`
}

type Students struct {
	Students []Student `xml:"student"`
}

func main() {
	file, _ := os.ReadFile("data.xml")
	var students Students
	xml.Unmarshal(file, &students)

	for _, s := range students.Students {
		fmt.Println("Name:", s.Name, "Marks:", s.Marks)
	}
}
