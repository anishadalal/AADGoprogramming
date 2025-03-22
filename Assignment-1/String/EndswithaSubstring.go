package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	fmt.Println(strings.HasSuffix(str, "world")) // true
}
