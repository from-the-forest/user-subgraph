package main

import (
	"log"
	"os"
	"user/graph"
	c "user/graph/context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	config := graph.Config{
		Resolvers: &graph.Resolver{},
	}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// create server
	r := gin.Default()

	// add middleware
	r.Use(c.ContextMiddleware())

	// add route handlers
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	// start server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Run()
}
