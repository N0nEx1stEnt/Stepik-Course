package main

import "fmt"

func main() {
	i := 1
	s := 0
	x := 0

	for i != 0 {
		fmt.Scan(&i)
		if i > x {
			x = i
			s = 1
		} else if i == x {
			s += 1
		}
	}
	fmt.Print(s)
}
