package main
import "fmt"

type user struct {
	Id             int
	Name, Location string
}
func (u *user) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}
type player struct {
	user
	GameId int
}
func main() {
	p := player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	fmt.Println(p.Greetings())
}