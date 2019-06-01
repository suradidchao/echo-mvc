package seed

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	userModel "github.com/suradidchao/echo-mvc/models/user"
	userAdapter "github.com/suradidchao/echo-mvc/models/user/adapters/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func SeedUsers(mongoConnectionURI string, db string) {
	names := []string{"Joe", "Jimmy", "William", "Jason", "Bill", "Thomas", "Isiah", "Vin"}
	nationality := []string{"TH", "MY", "US", "UK"}
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionURI))
	collection := mongoClient.Database(db).Collection("users")
	if err != nil {
		log.Fatal(err)
	}
	userMongoAdapter := userAdapter.NewAdapter(collection)
	userModel := userModel.NewModel(userMongoAdapter)
	user1 := map[string]interface{}{
		"name":        names[0],
		"age":         25,
		"nationality": nationality[2],
	}
	user2 := map[string]interface{}{
		"name":        names[2],
		"age":         30,
		"nationality": nationality[3],
	}

	_, err = userModel.CreateUser(user1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = userModel.CreateUser(user2)
	if err != nil {
		fmt.Println(err)
	}
}
