package main

import "fmt"

func main() {
	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 1+3, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}
}