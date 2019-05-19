package mongo_test

import (
	"testing"

	"github.com/suradidchao/echo-mvc/entities"
	"github.com/suradidchao/echo-mvc/models/user/adapters/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMapUserMongoToUser(t *testing.T) {
	userID := "5cd81b253810b40299cf91c2"
	objectID, _ := primitive.ObjectIDFromHex(userID)
	inputUser := mongo.UserMongo{
		ID:          objectID,
		Name:        "Jia",
		Age:         24,
		Nationality: "TH",
	}
	want := entities.User{
		ID:          userID,
		Name:        "Jia",
		Age:         24,
		Nationality: "TH",
	}
	got := mongo.MapUserMongoToUser(inputUser)
	if got != want {
		t.Errorf("MapUserMongo(%+v) = %+v; want %+v", inputUser, got, want)
	}
}
