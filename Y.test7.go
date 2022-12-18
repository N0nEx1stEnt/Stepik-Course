package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < 5; i++ {
		t := <-ticker.C
		fmt.Println(int(t.Sub(start).Seconds()))
	}
}
