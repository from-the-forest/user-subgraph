package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"user/graph/lib"
)

type UserRecord struct {
	ObjectID  primitive.ObjectID `bson: "_id" json: "_id"`
	ID        string             `bson: "id" json: "id"`
	FirstName string             `bson: "firstName" json: "firstName"`
	LastName  string             `bson: "lastName" json: "lastName"`
	Email     string             `bson: "email" json: "email"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: failed to load .env file")
	}

	userCollection, err := lib.GetUserCollection()
	if err != nil {
		panic(err)
	}

	const nodeId = "VXNlcjpkODBhOTNiZS00MGEwLTRhNTctODQ2YS1lZTU5MDY1ZmY1Mzc="

	result := userCollection.FindOne(context.Background(), bson.M{"id": bson.M{"$eq": nodeId}})
	if err != nil {
		panic(err)
	}

	var record UserRecord
	err = result.Decode(&record)
	if err != nil {
		panic(err)
	}

	fmt.Println(record)
}
