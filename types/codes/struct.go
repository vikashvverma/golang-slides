package main

import (
	"fmt"
	"time"
)

type BootCamp struct {
	// Latitude of the event
	Lat float64
	// Longitude of the event
	Lon float64
	// Date of the event
	Date time.Time
}

func main() {
	fmt.Println(BootCamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	})
}