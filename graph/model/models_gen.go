// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"user-subgraph/graph/scalar"
)

type Node interface {
	IsNode()
	// Relay node id
	GetID() string
}

type CreateUserInput struct {
	FirstName *string       `json:"firstName"`
	LastName  *string       `json:"lastName"`
	Email     *scalar.Email `json:"email"`
}

type PageInfo struct {
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	HasNextPage     bool    `json:"hasNextPage"`
}

type UpdateUserInput struct {
	ID        string        `json:"id"`
	FirstName *string       `json:"firstName"`
	LastName  *string       `json:"lastName"`
	Email     *scalar.Email `json:"email"`
}

// User type
type User struct {
	// User's first name
	FirstName string `json:"firstName"`
	// Node ID
	ID string `json:"id"`
	// User's last name
	LastName string `json:"lastName"`
	// User's full name (example of a compound field)
	FullName string `json:"fullName"`
	// User's email address
	Email scalar.Email `json:"email"`
}

func (User) IsNode() {}

// Relay node id
func (this User) GetID() string { return this.ID }

func (User) IsEntity() {}

type UsersConnection struct {
	PageInfo *PageInfo    `json:"pageInfo"`
	Edges    []*UsersEdge `json:"edges"`
}

type UsersEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
