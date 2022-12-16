package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timeStr := now.Format(time.RFC822) //Sun 19 Sep 2021 15:42:00 MSK
	fmt.Println(timeStr)
}
