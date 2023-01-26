package lib

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"user-subgraph/graph/model"
	"user-subgraph/graph/scalar"

	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByID(userCollection *mongo.Collection, id string) (*model.User, error) {
	userCollection, err := GetUserCollection()
	if err != nil {
		panic(err)
	}

	result := userCollection.FindOne(context.Background(), bson.M{"id": bson.M{"$eq": id}})
	if err != nil {
		panic(err)
	}

	var record UserRecord
	err = result.Decode(&record)
	if err != nil {
		panic(err)
	}

	user := &model.User{
		ID:        record.ID,
		FirstName: record.FirstName,
		LastName:  record.LastName,
		Email:     scalar.Email(record.Email),
	}

	return user, nil
}
