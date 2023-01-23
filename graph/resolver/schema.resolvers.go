package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"os"
	c "user/graph/context"
	graph "user/graph/generated"
	"user/graph/lib"
	"user/graph/model"
)

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	return id, nil
}

// Whoami is the resolver for the whoami field.
func (r *queryResolver) Whoami(ctx context.Context) (*model.User, error) {
	user := ctx.Value(c.UserCtxKey).(*model.User)
	return user, nil
}

// Host is the resolver for the host field.
func (r *queryResolver) Host(ctx context.Context) (string, error) {
	return os.Hostname()
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, input *model.UsersInput) (*model.UsersConnection, error) {
	// get nodes from (mocked) database
	users := lib.GetMockUsers()
	// node to edge
	edges := lib.Map(users, func(user model.User) *model.UsersEdge {
		return &model.UsersEdge{
			Node: &user,
		}
	})
	return &model.UsersConnection{
		PageInfo: &model.PageInfo{
			StartCursor:     nil,
			EndCursor:       nil,
			HasPreviousPage: false,
			HasNextPage:     false,
		},
		Edges: edges,
	}, nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	globalId, err := lib.FromGlobalId(id)
	if err != nil {
		return nil, fmt.Errorf("invalid global id %s", id)
	}
	switch globalId.Type {
	case "User":
		return lib.FindUserByID(globalId.ID)
	}
	return nil, fmt.Errorf("not implemented for type %s", globalId.Type)
}

// FullName is the resolver for the fullName field.
func (r *userResolver) FullName(ctx context.Context, obj *model.User) (string, error) {
	return obj.FirstName + " " + obj.LastName, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
