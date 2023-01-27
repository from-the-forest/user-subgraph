package lib

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"user-subgraph/graph/model"
)

func FindUserByID(userCollection *mongo.Collection, id string) (*model.User, error) {
	userCollection, err := GetUserCollection()
	if err != nil {
		return nil, err
	}

	result := userCollection.FindOne(context.Background(), bson.M{"id": bson.M{"$eq": id}})
	if err != nil {
		return nil, err
	}

	var record UserRecord
	err = result.Decode(&record)
	if err != nil {
		return nil, err
	}

	user := UserRecordToUserModel(record)

	return &user, nil
}
