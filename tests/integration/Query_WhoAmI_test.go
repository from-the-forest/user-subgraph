package main

import (
	"context"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func Test_Query_WhoAmI(t *testing.T) {

	t.Run("Query.whoami", func(t *testing.T) {
		graphqlEndpoint := os.Getenv("GRAPHQL_ENDPOINT")
		if graphqlEndpoint == "" {
			log.Fatal("you must specify a graphql endpoint to run the tests against")
		}

		// read query in from file - primarily for better developer experience writing the query
		file, err := os.ReadFile("./Query_WhoAmI.graphql")
		query := string(file)

		graphqlClient := graphql.NewClient(graphqlEndpoint)
		request := graphql.NewRequest(query)

		// set any variables
		//request.Var("key", "value")

		// set header fields
		//request.Header.Set("Cache-Control", "no-cache")

		// define a Context for the request
		ctx := context.Background()

		var response interface{}
		err = graphqlClient.Run(ctx, request, &response)
		if err != nil {
			panic(err)
		}

		assert.NotNil(t, response)
		assert.Nil(t, err)

		// TODO: would be nice to convert the response into some types to do assertions on
		//body, _ := json.Marshal(response)
		//fmt.Println(string(body))
	})
}
