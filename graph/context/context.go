package context

import (
	"context"
	"fmt"
	"os"
	"user/graph/model"

	"github.com/gin-gonic/gin"
)

type ContextKey struct {
	name string
}

var UserCtxKey = &ContextKey{"user"}
var EnvCtxkey = &ContextKey{"env"}

func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		fmt.Println(authorizationHeader)
		// TODO: parse jwt token to get user id to look up user record

		// user context value
		// TODO: get user out of request context
		user := &model.User{
			ID:        "1",
			FirstName: "Joe",
			LastName:  "Cuffney",
			Email:     "josephcuffney@gmail.com",
		}

		// set context values
		ctx := context.WithValue(c.Request.Context(), UserCtxKey, user)
		ctx = context.WithValue(ctx, EnvCtxkey, os.Getenv("ENV"))

		// add context to the request and proceed
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
