package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dannyoka/go-mongo/src/config"
	"github.com/dannyoka/go-mongo/src/repositories"
	"github.com/dannyoka/go-mongo/src/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string
	Value float64
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, sucka!")
	}
	config := config.New()
	mongoUri, err := config.Get("MONGO_URI")
	if err != nil {
		log.Fatalf("Error getting MONGO_URI: %v", err)
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()
	db := client.Database("test")
	collection := db.Collection("test")
	personRepository := repositories.NewPersonRepository(db)
	personService := services.NewPersonService(personRepository)
	person, err := personService.FindOneByName("pi")
	if err != nil {
		log.Fatalf("Error inserting document: %v", err)
	}
	fmt.Println(person.Name)
	fmt.Println(person.Value)
	fmt.Println(collection)
	fmt.Println("Hello World")
}
