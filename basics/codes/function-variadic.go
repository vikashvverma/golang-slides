package main

import "fmt"

func addNumbers(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(addNumbers(1,2,3))
}
