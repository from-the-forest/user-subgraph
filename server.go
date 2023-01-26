package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	c "user-subgraph/graph/context"
	generated "user-subgraph/graph/generated"
	"user-subgraph/graph/model"
	resolver "user-subgraph/graph/resolver"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	Production  string = "production"
	Development string = "development"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	config := generated.Config{
		Resolvers: &resolver.Resolver{},
	}

	// NOTE: this is only here because `Directives` is a struct not an interface :(
	// I'd much rather this be in `graph/directives` and included in the config similar to &graph.Resolver{},
	// https://github.com/99designs/gqlgen/issues/632
	config.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		// if !getCurrentUser(ctx).HasRole(role) {
		// 	// block calling the next resolver
		// 	return nil, fmt.Errorf("access denied")
		// }
		return nil, fmt.Errorf("access denied")
		// or let it pass through
		// return next(ctx)
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// c.Header("Access-Control-Allow-Origin", "https://fromtheforest.io")
		c.Header("Access-Control-Allow-Origin", "https://studio.apollographql.com")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Subgraph-Secret")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SubgraphSecretMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// []string -> string
		subgraphSecretHeader := strings.Join(c.Request.Header["X-Subgraph-Secret"], "")
		subgraphSecret := os.Getenv("SUBGRAPH_SECRET")
		// if the secret header is not present - the request is not authorized
		if subgraphSecretHeader != subgraphSecret {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: failed to load .env file")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Printf("Warning: no ENV specified - defaulting to %s", Production)
		env = Production
	}

	var ginMode = gin.ReleaseMode
	if env != Production {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// create server
	r := gin.Default()

	// add middleware
	r.Use(c.ContextMiddleware())
	if env == Production {
		r.Use(CorsMiddleware())
		r.Use(SubgraphSecretMiddleware())
	}

	// add route handlers
	r.GET("/", graphqlHandler())
	r.POST("/", graphqlHandler())
	if env != Production {
		r.GET("/playground", playgroundHandler())
	}

	// start server
	if env != Development {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

	r.Run()
}
