package main

import "fmt"

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
)

func main() {
	const Greeting = "नमस्ते"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
}