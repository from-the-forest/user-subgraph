package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	graph "user/graph/generated"
	"user/graph/lib"
	"user/graph/model"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	userCollection, err := lib.GetUserCollection()
	if err != nil {
		panic(err)
	}
	return lib.FindUserByID(userCollection, id)
}

// Entity returns graph.EntityResolver implementation.
func (r *Resolver) Entity() graph.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
