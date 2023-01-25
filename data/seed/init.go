package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"user/graph/lib"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"syreclabs.com/go/faker"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: failed to load .env file")
	}

	// connect to database
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnectionString))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())
	mongoDatabaseName := os.Getenv("MONGO_DATABASE")
	mongoUserCollectionName := os.Getenv("MONGO_USER_COLLECTION")
	userCollection := client.Database(mongoDatabaseName).Collection(mongoUserCollectionName)

	type UserRecord struct {
		ID        string
		FirstName string
		LastName  string
		Email     string
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
