package seed

import (
	"math/rand"
	"time"
)

// User is ...
type User struct {
	ID          int
	Name        string
	Age         int
	Nationality string
}

// GetUsers is ...
func GetUsers() []User {
	names := []string{"Joe", "Jimmy", "William", "Jason", "Bill", "Thomas", "Isiah", "Vin"}
	nationality := []string{"TH", "MY", "US", "UK"}
	users := []User{}
	rand.Seed(time.Now().Unix())
	for id, name := range names {
		newUser := User{
			ID:          id,
			Name:        name,
			Age:         rand.Intn(15) + 15,
			Nationality: nationality[rand.Intn(len(nationality))],
		}
		users = append(users, newUser)
	}
	return users
}
