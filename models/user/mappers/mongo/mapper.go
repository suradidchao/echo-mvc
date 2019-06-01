package mongo

import (
	"github.com/suradidchao/echo-mvc/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserMongo ...
type UserMongo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Age         int32              `bson:"age"`
	Nationality string             `bson:"nationality"`
}

// MapUserMongoToUser ...
func MapUserMongoToUser(user UserMongo) entities.User {
	return entities.User{
		ID:          user.ID.Hex(),
		Name:        user.Name,
		Age:         int(user.Age),
		Nationality: user.Nationality,
	}
}
