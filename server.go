package main

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	userController "github.com/suradidchao/echo-mvc/controllers/user"
	userModel "github.com/suradidchao/echo-mvc/models/user"
	userAdapter "github.com/suradidchao/echo-mvc/models/user/adapters/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27018"))
	collection := mongoClient.Database("db").Collection("users")
	if err != nil {
		log.Fatal(err)
	}

	userMongoAdapter := userAdapter.NewAdapter(collection)
	userModel := userModel.NewModel(userMongoAdapter)
	userController := userController.NewController(userModel)
	userController.RegisterRoutes("/users", e)

	e.Logger.Fatal(e.Start(":1323"))
}
