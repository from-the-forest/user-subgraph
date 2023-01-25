package context

import (
	"context"
	"log"
	"os"
	"user/graph/lib"

	"github.com/gin-gonic/gin"
)

type ContextKey struct {
	name string
}

var UserCtxKey = &ContextKey{"user"}
var EnvCtxkey = &ContextKey{"env"}
var UserCollectionCtxKey = &ContextKey{"userCollection"}

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

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	return func(c *gin.Context) {

		/**
		 * All code in this block is run once per request - keep that in mind.
		 * - nothing non-user specific (other things belong above)
		 * - nothing prohibatively expensive
		 */

		// TODO: get user out of request context
		authorizationHeader := c.GetHeader("Authorization")
		log.Printf("Authorization Header: %s", authorizationHeader)
		// TODO: parse jwt token to get user id to look up user record
		userId := "VXNlcjpiNTNhMzViMC03MmJmLTQ2OWItOGE1OS01YTI3Zjg0Nzc2Mzk"
		user, err := lib.FindUserByID(userCollection, userId)
		if err != nil {
			log.Fatal("Failed to find user by id")
		}

		// ////////////////////////////////////////////////////////////////////////
		// set context values, and add them to the request
		// ////////////////////////////////////////////////////////////////////////

		ctx := context.WithValue(c.Request.Context(), UserCtxKey, user)
		ctx = context.WithValue(ctx, EnvCtxkey, env)
		ctx = context.WithValue(ctx, UserCollectionCtxKey, userCollection)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
