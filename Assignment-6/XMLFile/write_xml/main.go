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

func main() {
	students := []Student{
		{"Alice", 92},
		{"Bob", 87},
	}
	data, _ := xml.MarshalIndent(students, "", "  ")
	os.WriteFile("output.xml", data, 0644)
	fmt.Println("XML file written successfully.")
}
