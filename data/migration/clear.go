package main

import (
	"context"
	"fmt"
	"user/graph/lib"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	userCollection, err := lib.GetUserCollection()
	if err != nil {
		panic(err)
	}

	result, err := userCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
