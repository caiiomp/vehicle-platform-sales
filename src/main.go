package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		mongoURI = os.Getenv("MONGO_URI")
		// mongoDatabase = os.Getenv("MONGO_DATABASE")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("could not initialize mongodb client: %v", err)
	}

	if err = mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Collections
	// vehiclesCollection := mongoClient.Database(mongoDatabase).Collection("vehicles")
	// salesCollection := mongoClient.Database(mongoDatabase).Collection("sales")

}
