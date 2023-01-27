package context

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"user-subgraph/graph/model"
)

// NOTE: we use these getters for getting context to ensure we retain type safety
// https://www.calhoun.io/pitfalls-of-context-values-and-how-to-avoid-or-mitigate-them/

func GetUser(ctx context.Context) *model.User {
	user, ok := ctx.Value(UserCtxKey).(*model.User)
	if !ok {
		log.Fatal(ok)
		// TODO: Log this issue
		return nil
	}
	return user
}

func GetUserCollection(ctx context.Context) *mongo.Collection {
	userCollection, ok := ctx.Value(UserCollectionCtxKey).(*mongo.Collection)
	if !ok {
		// Log this issue
		return nil
	}
	return userCollection
}
