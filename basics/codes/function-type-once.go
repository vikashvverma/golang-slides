package main

import "fmt"

func addWithTypeOnce(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(addWithTypeOnce(42, 13))
}