package main

import (
	"context"
	"fmt"
	"user/graph/lib"

	"github.com/google/uuid"
	"syreclabs.com/go/faker"
)

type UserRecord struct {
	ID        string `bson:"_id"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
}

func main() {
	userCollection, err := lib.GetUserCollection()
	if err != nil {
		panic(err)
	}

	NUM_USERS := 20
	users := make([]interface{}, 0)
	for i := 1; i <= NUM_USERS; i++ {
		users = append(users, UserRecord{
			ID:        lib.ToGlobalID("User", uuid.NewString()),
			FirstName: faker.Name().FirstName(),
			LastName:  faker.Name().LastName(),
			Email:     faker.Internet().Email(),
		})
	}
	fmt.Println(users)

	result, err := userCollection.InsertMany(context.Background(), users)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
