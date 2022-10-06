package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	if i == 0 {
		fmt.Printf("Ноль")
	} else if i > 0 {
		fmt.Printf("Число положительное")
	} else {
		fmt.Printf("Число отрицательное")
	}
	return
}
