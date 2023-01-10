package lib

import (
	"user/graph/model"
)

func FindUserByID(id string) (*model.User, error) {
	user := &model.User{
		ID:        "1",
		FirstName: "Joe",
		LastName:  "Cuffney",
		Email:     "josephcuffney@gmail.com",
	}
	return user, nil
}
