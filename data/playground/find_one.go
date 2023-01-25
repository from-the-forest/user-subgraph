package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"user/graph/lib"
)

func main() {
	userCollection, err := lib.GetUserCollection()
	if err != nil {
		panic(err)
	}

	result := userCollection.FindOne(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
