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
	file, err := os.ReadFile("data.xml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var students Students
	err = xml.Unmarshal(file, &students)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	for _, s := range students.Students {
		fmt.Println("Name:", s.Name, "Marks:", s.Marks)
	}
}
