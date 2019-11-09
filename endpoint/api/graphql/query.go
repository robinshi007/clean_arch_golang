package graphql

import (
	"github.com/graphql-go/graphql"

	"clean_arch/endpoint/api/graphql/field"
	"clean_arch/infra"
)

// NewRootQuery -
func NewRootQuery(db infra.DB) *graphql.Object {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"userList": field.NewUserListField(db),
		},
	})
	return rootQuery
}
