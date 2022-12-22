package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t stardater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour())
	return 1000 + day + h
}

func main() {
	day := time.Date(2022, 12, 22, 17, 27, 14, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))
}
