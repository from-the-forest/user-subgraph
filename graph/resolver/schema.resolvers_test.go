package resolver

import (
	"context"
	"fmt"
	"testing"
	c "user/graph/context"
	"user/graph/model"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	t.Run("Query.whoami", func(t *testing.T) {
		ctx := context.Background()
		user := &model.User{
			ID:        "1",
			FirstName: "some",
			LastName:  "user",
		}
		ctx = context.WithValue(ctx, c.UserCtxKey, user)
		r := Resolver{}
		resp, err := r.Query().Whoami(ctx)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
		fmt.Println(resp)
		assert.Equal(t, user, resp)
	})
}
