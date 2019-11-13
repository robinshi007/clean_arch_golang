package graphql

import (
	"github.com/graphql-go/graphql"

	"clean_arch/endpoint/api/graphql/field"
)

// NewRootMutation -
func NewRootMutation() *graphql.Object {
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": field.NewCreateUserField(),
		},
	})
	return rootMutation
}
