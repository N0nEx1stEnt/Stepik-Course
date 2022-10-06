package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
}
