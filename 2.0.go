package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	bytes, err := json.Marshal(locations)
	exitOnError(err)

	fmt.Println(string(bytes))
}
