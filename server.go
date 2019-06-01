package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	userController "github.com/suradidchao/echo-mvc/controllers/user"
	userModel "github.com/suradidchao/echo-mvc/models/user"
	userAdapter "github.com/suradidchao/echo-mvc/models/user/adapters/mongo"
	"github.com/suradidchao/echo-mvc/seed"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	env := flag.String("env", "local", "Input either local or production")

	var mongoConnectionURI string
	var db string
	flag.Parse()
	if *env == "local" {
		mongoConnectionURI = "localhost:27017"
		db = "db"
		seed.SeedUsers(mongoConnectionURI, db)

	} else {
		mongoConnectionURI = "mongodb://root:root1!@ds157946.mlab.com:57946/echo-mvc"
		db = "echo-mvc"
	}
	log.Printf("env: %s\n", *env)
	log.Printf("Mongo connection uri: %s\n", mongoConnectionURI)
	log.Printf("Mongo default db: %s\n", db)
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionURI))
	collection := mongoClient.Database(db).Collection("users")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userMongoAdapter := userAdapter.NewAdapter(collection)
	userModel := userModel.NewModel(userMongoAdapter)
	userController := userController.NewController(userModel)
	userController.RegisterRoutes("/users", e)

	e.Logger.Fatal(e.Start(":1323"))
}
