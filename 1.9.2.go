package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)

	c := i % 1000
	i = i / 1000
	c = c%10 + c/100 + c/10%10
	i = i%10 + i/100 + i/10%10
	if c == i {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
}
