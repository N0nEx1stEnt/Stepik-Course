package main

import "fmt"

type celsius float64

type location struct {
	lat, long float64
}

type temperature struct {
	high, low celsius
}

type report struct {
	day         int
	temperature temperature
	location    location
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{day: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)

	fmt.Printf("a balmy %v C\n", report.temperature.high)

	fmt.Printf("average temperature on Mars %v C\n", t.average())

	fmt.Printf("average temperature var2.0 %v C\n", report.temperature.average())

}
