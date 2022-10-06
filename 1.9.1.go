package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	if i == 10000 {
		fmt.Print(1)
	} else if i > 1000 {
		fmt.Print(i / 1000)
	} else if i > 100 {
		fmt.Print(i / 100)
	} else if i > 10 {
		fmt.Print(i / 10)
	} else {
		fmt.Print(i)
	}
}
