package main

import (
	"fmt"
)

type bootcamp struct {
	Lat float64
	Lon float64
}

func main() {
	x := new(bootcamp)
	y := &bootcamp{}
	fmt.Println(*x == *y)
}