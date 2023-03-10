package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	c "user-subgraph/graph/context"
	graph1 "user-subgraph/graph/generated"
	"user-subgraph/graph/lib"
	"user-subgraph/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgryski/trifles/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	newRecord := lib.UserRecord{
		ID:        lib.ToGlobalID("User", uuid.UUIDv4()),
		FirstName: *input.FirstName,
		LastName:  *input.LastName,
		Email:     string(*input.Email),
	}
	userCollection := c.GetUserCollection(ctx)
	_, err := userCollection.InsertOne(context.Background(), newRecord)
	if err != nil {
		graphql.AddError(ctx, err)
	}
	user := lib.UserRecordToUserModel(newRecord)
	return &user, err
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	userCollection := c.GetUserCollection(ctx)
	currentUser, err := lib.FindUserByID(userCollection, input.ID)
	//updatesToUser := model.User{
	//	FirstName: *input.FirstName,
	//	LastName:  *input.LastName,
	//	Email:     *input.Email,
	//}
	//user := lib.Merge(currentUser, updatesToUser).(model.User)
	return currentUser, err
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	userCollection := c.GetUserCollection(ctx)
	user, err := lib.FindUserByID(userCollection, id)
	if err != nil {
		graphql.AddError(ctx, err)
	}
	userCollection.DeleteOne(context.Background(), bson.M{"id": id})
	return user, nil
}

// Whoami is the resolver for the whoami field.
func (r *queryResolver) Whoami(ctx context.Context) (*model.User, error) {
	//// TODO: remove this - it's just an example of using the gql client within a resolver
	//gqlClient := c.GetGraphQLClient(ctx)
	//query := `{
	//	whoami { id }
	//}`
	//request := gql.NewRequest(query)
	//var response interface{}
	//err := gqlClient.Run(ctx, request, &response)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//body, _ := json.Marshal(response)
	//fmt.Println(string(body))

	user := c.GetUser(ctx)
	return user, nil
}

// SearchUsers is the resolver for the searchUsers field.
func (r *queryResolver) SearchUsers(ctx context.Context, first *int, after *string) (*model.UsersConnection, error) {
	userCollection := c.GetUserCollection(ctx)

	// slice to hold users
	users := make([]model.User, 0)

	// increment "first" by 1 so that we can determine if there is a next page
	limit := int64(*first) + int64(1)
	options := &options.FindOptions{
		// NOTE: might make more sense for the id portion of the node id to be `_id` but, I am going to try to have it work on a secondary index
		// specify index name
		Hint:  "by-id",
		Limit: &limit,
		// TODO: do I need a sort key here?
		//Sort: bson.D{{"id", 1}},
		// do not include the _id field in the response
		Projection: bson.M{"_id": 0},
	}
	// use `bson.D{}` as opposed to `bson.M{}` because order matters for pagination
	filter := bson.D{}
	if after != nil {
		filter = bson.D{{Key: "id", Value: bson.D{{Key: "$gte", Value: *after}}}}
	}
	cursor, err := userCollection.Find(context.Background(), filter, options)
	if err != nil {
		graphql.AddError(ctx, err)
	}

	var userRecords []lib.UserRecord
	err = cursor.All(context.Background(), &userRecords)
	if err != nil {
		graphql.AddError(ctx, err)
	}

	for _, userRecord := range userRecords {
		user := lib.UserRecordToUserModel(userRecord)
		users = append(users, user)
	}

	// get the pageInfo while we have `first + 1` elements, so we can get the endCursor (which is exclusive)
	startCursor := users[0].ID
	endCursor := users[len(users)-1].ID
	hasNextPage := len(users) > *first
	hasPreviousPage := after != nil
	// remove the extra element we fetched before building the edges
	users = users[:len(users)-1]

	// node to edge
	edges := lib.Map(users, func(user model.User) *model.UsersEdge {
		return &model.UsersEdge{
			Cursor: user.ID,
			Node:   &user,
		}
	})

	return &model.UsersConnection{
		PageInfo: &model.PageInfo{
			StartCursor:     &startCursor,
			EndCursor:       &endCursor,
			HasPreviousPage: hasPreviousPage,
			HasNextPage:     hasNextPage,
		},
		Edges: edges,
	}, nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	userCollection := c.GetUserCollection(ctx)
	globalId, err := lib.FromGlobalId(id)
	if err != nil {
		return nil, fmt.Errorf("invalid global id %s", id)
	}
	switch globalId.Type {
	case "User":
		return lib.FindUserByID(userCollection, id)
	}
	// TODO: how would this type `Node` be resolved from a single service without being aware of other types
	// i.e from relay-subgraph which has no awareness of user-subgraph - you wouldn't have a method to call?
	return nil, fmt.Errorf("not implemented for type %s", globalId.Type)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]model.Node, error) {
	for _, id := range ids {
		// 1. group by type
		_, err := lib.FromGlobalId(id)
		if err != nil {
			return nil, fmt.Errorf("invalid global id %s", id)
		}
	}
	graphql.AddError(ctx, fmt.Errorf("not implemented"))
	// 2. fetch nodes for a respective type
	var nodes []model.Node = []model.Node{}
	return nodes, nil
}

// FullName is the resolver for the fullName field.
func (r *userResolver) FullName(ctx context.Context, obj *model.User) (string, error) {
	return obj.FirstName + " " + obj.LastName, nil
}

// Mutation returns graph1.MutationResolver implementation.
func (r *Resolver) Mutation() graph1.MutationResolver { return &mutationResolver{r} }

// Query returns graph1.QueryResolver implementation.
func (r *Resolver) Query() graph1.QueryResolver { return &queryResolver{r} }

// User returns graph1.UserResolver implementation.
func (r *Resolver) User() graph1.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
