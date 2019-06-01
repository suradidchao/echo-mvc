package mongo

import (
	"context"
	"time"

	"github.com/suradidchao/echo-mvc/entities"
	mapper "github.com/suradidchao/echo-mvc/models/user/mappers/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Adapter ...
type Adapter struct {
	Collection *mongo.Collection
}

// NewAdapter ...
func NewAdapter(collection *mongo.Collection) Adapter {
	return Adapter{
		Collection: collection,
	}
}

// FindMany ...
func (a Adapter) FindMany() ([]entities.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	cur, err := a.Collection.Find(ctx, bson.D{})
	if err != nil {
		return []entities.User{}, err
	}

	defer cur.Close(ctx)
	var users []entities.User
	for cur.Next(ctx) {
		var user mapper.UserMongo

		err := cur.Decode(&user)
		if err != nil {
			return []entities.User{}, err
		}

		currentUser := mapper.MapUserMongoToUser(user)
		users = append(users, currentUser)
	}
	if err := cur.Err(); err != nil {
		return []entities.User{}, err
	}
	return users, nil
}

// FindOne ...
func (a Adapter) FindOne(userID string) (entities.User, error) {
	var user mapper.UserMongo
	var foundUser entities.User
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return foundUser, err
	}
	filter := bson.M{"_id": objectID}
	err = a.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return foundUser, err
	}
	foundUser = mapper.MapUserMongoToUser(user)
	return foundUser, nil
}

// CreateOne ...
func (a Adapter) CreateOne(userData map[string]interface{}) (entities.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	data := bson.M{
		"name":        userData["name"],
		"age":         userData["age"],
		"nationality": userData["nationality"],
	}
	res, err := a.Collection.InsertOne(ctx, data)
	if err != nil {
		return entities.User{}, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	user := entities.User{
		ID:          id.Hex(),
		Name:        userData["name"].(string),
		Age:         userData["age"].(int),
		Nationality: userData["nationality"].(string),
	}
	return user, nil
}

// UpdateOne ...
func (a Adapter) UpdateOne(userID string, userData map[string]interface{}) (entities.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	data := bson.M{
		"name":        userData["name"],
		"age":         userData["age"],
		"nationality": userData["nationality"],
	}
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return entities.User{}, err
	}

	_, err = a.Collection.UpdateOne(
		ctx,
		bson.D{
			{"_id", objectID},
		},
		bson.D{
			{"$set", bson.D{
				{"name", data["name"]},
				{"age", data["age"]},
				{"nationality", data["nationality"]},
			}},
		},
	)
	if err != nil {
		return entities.User{}, err
	}
	updatedUser := entities.User{
		ID:          userID,
		Name:        userData["name"].(string),
		Age:         userData["age"].(int),
		Nationality: userData["nationality"].(string),
	}
	return updatedUser, nil
}

// DeleteOne ...
func (a Adapter) DeleteOne(userID string) (entities.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	ObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return entities.User{}, err
	}

	deletingUser, err := a.FindOne(userID)
	if err != nil {
		return entities.User{}, err
	}

	_, err = a.Collection.DeleteOne(
		ctx,
		bson.D{
			{"_id", ObjectID},
		},
	)
	if err != nil {
		return entities.User{}, err
	}
	return deletingUser, nil
}
