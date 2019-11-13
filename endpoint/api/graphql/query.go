package graphql

import (
	"github.com/graphql-go/graphql"

	"clean_arch/endpoint/api/graphql/field"
)

// NewRootQuery -
func NewRootQuery() *graphql.Object {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"userList": field.NewUserListField(),
		},
	})
	return rootQuery
}
