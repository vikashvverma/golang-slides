package main
import "fmt"
func main() {
	numbers := []int{4, 6, 8} // Initializing a slice
	mySlice := make([]int, 6) // Making a slice

	// Assigning values
	mySlice[0] = 2
	mySlice[1] = 3
	mySlice[2] = 5
	mySlice[3] = 7
	mySlice = append(mySlice, 7) // Appending to slice
	mySlice = append(mySlice, 11, 13) // Appending more than one value
	mySlice = append(mySlice, numbers...) // Appending using ellipses

	fmt.Println(mySlice)

	fmt.Println(mySlice[1:4])

	fmt.Println(mySlice[:3]) // missing low index implies 0

	fmt.Println(mySlice[4:]) // missing high index implies len(s)

	fmt.Println(len(mySlice)) // Length of a slice
}