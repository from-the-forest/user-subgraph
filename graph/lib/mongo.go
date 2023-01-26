package lib

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRecord struct {
	ObjectID  primitive.ObjectID `bson: "_id" json: "_id"`
	ID        string             `bson: "id" json: "id"`
	FirstName string             `bson: "firstName" json: "firstName"`
	LastName  string             `bson: "lastName" json: "lastName"`
	Email     string             `bson: "email" json: "email"`
}

func GetUserCollection() (*mongo.Collection, error) {
	// connect to database
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnectionString))
	if err != nil {
		return nil, err
	}
	mongoDatabaseName := os.Getenv("MONGO_DATABASE")
	mongoUserCollectionName := os.Getenv("MONGO_USER_COLLECTION")
	userCollection := client.Database(mongoDatabaseName).Collection(mongoUserCollectionName)

	return userCollection, nil
}
