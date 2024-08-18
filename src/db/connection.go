package db

import (
	"context"
	"log"

	"github.com/dannyoka/go-mongo/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(config *config.Config) *mongo.Client {
	mongoUri := config.Get("MONGO_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")
	return client
}

func InitDB(config *config.Config, dbName string) *mongo.Database {
	conn := Connect(config)
	return conn.Database(dbName)
}
