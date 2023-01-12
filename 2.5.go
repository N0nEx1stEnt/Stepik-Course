package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello"
	fmt.Println(string(s[2]))
	strings.Replace(s, string(s[1]), "", 2)
	fmt.Println(s)
}
