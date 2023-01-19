package lib

import (
	"user/graph/model"
)

func FindUserByID(id string) (*model.User, error) {
	user := &model.User{
		ID:        "1",
		FirstName: "Nox",
		LastName:  "Cuffney",
		Email:     "noxman@gmail.com",
	}
	return user, nil
}
