package main

import (
	"fmt"
)

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	year := 0

	for {
		x = x + (x*p)/100
		year++
		if x >= y {
			fmt.Print(year)
			break
		}
	}
}
