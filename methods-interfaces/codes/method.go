package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := User{"Black", "Panther"}
	fmt.Println(u.Greeting())
}