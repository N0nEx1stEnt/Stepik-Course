package main

import (
	"fmt"
)

func main() {
	var c, i int
	fmt.Scan(&c, &i)
	var s int

	for x := c; x <= i; x++ {
		s += x
	}
	fmt.Print(s)
}
