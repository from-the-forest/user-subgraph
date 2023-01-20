package main

import (
	"log"
	"os"
	"strings"
	"user/graph"
	c "user/graph/context"
	generated "user/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	config := generated.Config{
		Resolvers: &graph.Resolver{},
	}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/v1/graphql/user")

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

		// TODO: this should be in an auth middleware
		// []string -> string
		subgraphSecretHeader := strings.Join(c.Request.Header["X-Subgraph-Secret"], "")
		subgraphSecret := os.Getenv("SUBGRAPH_SECRET")
		// if the secret header is not present - the request is not authorized
		if subgraphSecretHeader != subgraphSecret {
			c.AbortWithStatus(401)
			return
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// create server
	r := gin.Default()

	// add middleware
	r.Use(c.ContextMiddleware())
	r.Use(CorsMiddleware())

	// add route handlers
	r.GET("/v1/graphql/user", graphqlHandler())
	r.POST("/v1/graphql/user", graphqlHandler())
	r.GET("/v1/graphql/user/playground", playgroundHandler())

	// start server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Run()
}
