package lib

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
