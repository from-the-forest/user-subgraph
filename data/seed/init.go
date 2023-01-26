package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"

	"user-subgraph/graph/lib"

	"github.com/google/uuid"
	"syreclabs.com/go/faker"
)

type UserRecord struct {
	ID        string `bson:"id"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
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

	const numUsers = 20
	users := make([]interface{}, 0)
	for i := 1; i <= numUsers; i++ {
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
