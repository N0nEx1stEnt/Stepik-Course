package main

import (
	"fmt"
)

func main() {
	var i, x int
	fmt.Scan(&i)
	s := 0

	for j := 0; j < i; j++ {
		fmt.Scan(&x)
		if (x/10 < 10 && x/10 > 0) && x%8 == 0 {
			s += x
		}
	}
	fmt.Print(s)
}
