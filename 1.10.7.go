package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	sum := 0
	for z := 0; z < 4; z++ {
		if x == 0 {
			break
		}
		i := x % 10
		x /= 10
		l := y
		for k := 0; k < 4; k++ {
			j := l % 10
			l /= 10
			if i == j {
				sum = (sum + i) * 10
				continue
			}
		}
	}
	sum /= 10
	for sum > 0 {
		i := sum % 10
		sum /= 10
		fmt.Print(i, " ")
	}
}
