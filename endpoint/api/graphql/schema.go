package graphql

import (
	"github.com/graphql-go/graphql"

	"clean_arch/infra/util"
)

// NewSchema -
func NewSchema() graphql.Schema {
	var schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    NewRootQuery(),
		Mutation: NewRootMutation(),
	})

	util.FailedIf(err)

	return schema
}
