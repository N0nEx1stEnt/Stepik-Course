package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	c := i / 100
	if c == i/10%10 {
		fmt.Printf("NO")
	} else if c == i%10 {
		fmt.Printf("NO")
	} else if i/10%10 == i%10 {
		fmt.Printf("NO")
	} else {
		fmt.Printf("YES")
	}
}
