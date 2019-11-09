package graphql

import (
	"github.com/graphql-go/graphql"

	"clean_arch/endpoint/api/graphql/field"
	"clean_arch/infra"
)

// NewRootMutation -
func NewRootMutation(db infra.DB) *graphql.Object {
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": field.NewCreateUserField(db),
		},
	})
	return rootMutation
}
