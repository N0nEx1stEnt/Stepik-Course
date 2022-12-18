package main

import "fmt"

type celsius float64

type temperature struct {
	high, low celsius
}

type day int

type location struct {
	lat, long float64
}

type report struct {
	day
	temperature
	location
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (l location) days(l2 location) int {
	return 5
}

func main() {
	report := report{
		day:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -71.0},
	}
	fmt.Printf("average temperature on Mars %v C\n", report.average())

	fmt.Printf("High temperature %v C\n", report.high)

	fmt.Printf("Lat of the transport %v \n", report.lat)
}
