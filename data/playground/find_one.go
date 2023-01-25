package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"user/graph/lib"
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

	result := userCollection.FindOne(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
