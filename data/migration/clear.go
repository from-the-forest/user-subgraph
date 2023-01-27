package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"user-subgraph/graph/lib"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: failed to load .env file")
	}

	userCollection, err := lib.GetUserCollection()
	if err != nil {
		panic(err)
	}

	result, err := userCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted: %v records\n", result.DeletedCount)
}
