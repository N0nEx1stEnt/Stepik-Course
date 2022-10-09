package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	i := 0

	for i <= n {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)
			break
		}
		i++
	}
}
