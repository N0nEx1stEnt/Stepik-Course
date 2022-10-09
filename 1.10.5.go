package main

import (
	"fmt"
)

func main() {
	for {
		var i int
		fmt.Scan(&i)
		if i < 10 {
			continue
		} else if i > 100 {
			break
		}
		fmt.Println(i)
	}
}
