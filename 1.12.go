package main

import (
	"fmt"
)

func main() {
	var workArray [10]uint8
	for i := 0; i < 10; i++ {
		var x uint8
		fmt.Scan(&x)
		workArray[i] = x
	}

	for i := 0; i < 3; i++ {
		var x, y uint8
		fmt.Scan(&x, &y)

		z := workArray[x]
		workArray[x] = workArray[y]
		workArray[y] = z
	}

	for i := 0; i < len(workArray); i++ {
		fmt.Printf("%v ", workArray[i])
	}
}
