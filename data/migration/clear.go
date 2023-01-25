package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: failed to load .env file")
	}

	// connect to database
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnectionString))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())
	mongoDatabaseName := os.Getenv("MONGO_DATABASE")
	mongoUserCollectionName := os.Getenv("MONGO_USER_COLLECTION")
	userCollection := client.Database(mongoDatabaseName).Collection(mongoUserCollectionName)

	result, err := userCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
