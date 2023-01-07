package context

import (
	"context"
	"user/graph/model"

	"github.com/gin-gonic/gin"
)

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{"user"}
var EnvCtxkey = &contextKey{"env"}

func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// user context value
		// TODO: get user out of request context
		user := &model.User{
			ID:        "1",
			FirstName: "Joe",
			LastName:  "Cuffney",
		}

		// set context values
		ctx := context.WithValue(c.Request.Context(), UserCtxKey, user)
		ctx = context.WithValue(ctx, EnvCtxkey, "dev")

		// add context to the request and proceed
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
