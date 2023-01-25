package lib

import (
	"user/graph/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByID(userCollection *mongo.Collection, id string) (*model.User, error) {
	user := &model.User{
		ID:        ToGlobalID("User", "1"),
		FirstName: "Willy",
		LastName:  "Cuffney",
		Email:     "littlewilly@gmail.com",
	}
	return user, nil
}
