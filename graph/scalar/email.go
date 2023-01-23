package scalar

import (
	"context"
	"fmt"
	"io"
	"net/mail"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Email string

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// for incoming email values (i.e from a client on an input type)
func (e *Email) UnmarshalGQLContext(ctx context.Context, v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Email must be a string")
	}

	if !isValidEmail(str) {
		return fmt.Errorf("%s is not a valid email", str)
	}

	*e = Email(str)

	return nil
}

// for outgoing email values (i.e returned on a type)
func (e Email) MarshalGQLContext(ctx context.Context, w io.Writer) error {
	str := string(e)

	if !isValidEmail(str) {
		// NOTE: this adds an error to the gql response without preventing data from
		// being returned.
		graphql.AddErrorf(ctx, "%s is not a valid email", str)
	}

	w.Write([]byte(strconv.Quote(string(e))))

	return nil
}
