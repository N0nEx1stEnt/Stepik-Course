package main

import (
	"fmt"
)

func main() {

	var r float64
	fmt.Scan(&r)

	if r <= 0 {
		fmt.Printf("число %2.2f не подходит", r)
	} else if r > 10000 {
		fmt.Printf("%e", r)
	} else {
		fmt.Printf("%.4f", r*r)
	}
}
