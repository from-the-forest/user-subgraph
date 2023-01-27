package context

import (
	"context"
	"log"
	"os"
	"user-subgraph/graph/lib"

	"github.com/gin-gonic/gin"
)

type ContextKey struct {
	name string
}

var UserCtxKey = &ContextKey{"user"}
var EnvCtxkey = &ContextKey{"env"}
var UserCollectionCtxKey = &ContextKey{"userCollection"}
var GraphQLCtxKey = &ContextKey{"graphql"}

func ContextMiddleware() gin.HandlerFunc {
	/**
	 * All code before the return statement is run at server startup - not per
	 * request.  As such anything that would be "expensive" to do every request
	 * should go in this section. i.e non-user specific stuff
	 */

	// ////////////////////////////////////////////////////////////////////////
	// MongoDB Collections
	// ////////////////////////////////////////////////////////////////////////

	userCollection, err := lib.GetUserCollection()
	if err != nil {
		panic(err)
	}

	// ////////////////////////////////////////////////////////////////////////
	// Environment
	// ////////////////////////////////////////////////////////////////////////

	env := os.Getenv("ENV")

	// ////////////////////////////////////////////////////////////////////////
	// GraphQL.  (make supergraph calls from within a subgraph)
	// ////////////////////////////////////////////////////////////////////////

	// TODO: create a function to execute graphql queries that will be available to resolvers
	graphql := 0

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return func(c *gin.Context) {

		/**
		 * All code in this block is run once per request - keep that in mind.
		 * - nothing non-user specific (other things belong above)
		 * - nothing prohibitively expensive
		 */

		// TODO: get user out of request context
		authorizationHeader := c.GetHeader("Authorization")
		log.Printf("Authorization Header: %s", authorizationHeader)
		// TODO: parse jwt token to get user id and look up user record
		userId := "VXNlcjowODEyMzE3ZC1hYzE0LTRkODktOTMwZi03MDgyMmZjNzdjMGI="
		user, err := lib.FindUserByID(userCollection, userId)
		if err != nil {
			panic(err)
		}

		// ////////////////////////////////////////////////////////////////////////
		// set context values, and add them to the request
		// ////////////////////////////////////////////////////////////////////////
		ctx := context.WithValue(c.Request.Context(), UserCtxKey, user)
		ctx = context.WithValue(ctx, EnvCtxkey, env)
		ctx = context.WithValue(ctx, UserCollectionCtxKey, userCollection)
		ctx = context.WithValue(ctx, GraphQLCtxKey, graphql)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
