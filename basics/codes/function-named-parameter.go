package main

import "fmt"

func locationWithNamedReturns(name, city string) (region, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}
	return
}

func main() {
	region, continent := locationWithNamedReturns("Matt", "LA")
	fmt.Printf("%s lives in %s", region, continent)
}
