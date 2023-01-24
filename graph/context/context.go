package context

import (
	"context"
	"fmt"
	"log"
	"user/graph/lib"

	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContextKey struct {
	name string
}

var UserCtxKey = &ContextKey{"user"}
var EnvCtxkey = &ContextKey{"env"}
var UserCollectionCtxKey = &ContextKey{"userCollection"}

func ContextMiddleware() gin.HandlerFunc {

	fmt.Println("context: run at server startup - not rerun per request (non user specific stuff goes here)")

	// ////////////////////////////////////////////////////////////////////////
	// MongoDB
	// ////////////////////////////////////////////////////////////////////////

	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnectionString))
	if err != nil {
		panic(err)
	}
	mongoDatabaseName := os.Getenv("MONGO_DATABASE")
	mongoUserCollectionName := os.Getenv("MONGO_USER_COLLECTION")
	userCollection := client.Database(mongoDatabaseName).Collection(mongoUserCollectionName)

	// ////////////////////////////////////////////////////////////////////////
	// Environment
	// ////////////////////////////////////////////////////////////////////////

	env := os.Getenv("ENV")

	return func(c *gin.Context) {
		fmt.Println("context: run once per request (user specific stuff goes here")

		// TODO: get user out of request context
		authorizationHeader := c.GetHeader("Authorization")
		log.Printf("Authorization Header: %s", authorizationHeader)
		// TODO: parse jwt token to get user id to look up user record
		userId := "1"
		user, err := lib.FindUserByID(userId)
		if err != nil {
			log.Fatal("Failed to find user by id")
		}

		// set context values
		ctx := context.WithValue(c.Request.Context(), UserCtxKey, user)
		ctx = context.WithValue(ctx, EnvCtxkey, env)
		ctx = context.WithValue(ctx, UserCollectionCtxKey, userCollection)

		// add context to the request and proceed
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
