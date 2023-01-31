package context

import (
	"context"
	"github.com/machinebox/graphql"
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
		return nil
	}
	return user
}

func GetUserCollection(ctx context.Context) *mongo.Collection {
	userCollection, ok := ctx.Value(UserCollectionCtxKey).(*mongo.Collection)
	if !ok {
		log.Fatal(ok)
		return nil
	}
	return userCollection
}

func GetGraphQLClient(ctx context.Context) *graphql.Client {
	gqlClient, ok := ctx.Value(GraphQLCtxKey).(*graphql.Client)
	if !ok {
		log.Fatal(ok)
		return nil
	}
	return gqlClient
}
