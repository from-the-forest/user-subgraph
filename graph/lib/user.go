package lib

import (
	"user/graph/model"
)

func FindUserByID(id string) (*model.User, error) {
	user := &model.User{
		ID:        ToGlobalID("User", "1"),
		FirstName: "Willy",
		LastName:  "Cuffney",
		Email:     "littlewilly@gmail.com",
	}
	return user, nil
}

// temporary

func GetMockUsers() []model.User {

	users := []model.User{
		{
			ID:        ToGlobalID("User", "1"),
			FirstName: "Nox",
			LastName:  "Cuffney",
			Email:     "noxman@dog.com",
		},
		{
			ID:        ToGlobalID("User", "2"),
			FirstName: "Willy",
			LastName:  "Cuffney",
			Email:     "littlewilly@dog.com",
		},
	}
	return users
}
