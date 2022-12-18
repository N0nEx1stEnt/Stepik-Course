package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

func (l location) description() string {
	l2 := fmt.Sprintf("%.1f", l.lat)
	l3 := fmt.Sprintf("%.1f", l.long)
	s2 := l.name + " " + l2 + " " + l3
	return s2
}

type gps struct {
	world       world
	current     location
	destination location
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.destination.description())
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{"Bradbury Landing", -4.5895, 137.4417}
	elysium := location{"Elysium Planitia", 4.5, 135.9}

	gps := gps{
		world:       mars,
		current:     bradbury,
		destination: elysium,
	}

	curiosity := rover{
		gps: gps,
	}

	fmt.Println(curiosity.message())
}
