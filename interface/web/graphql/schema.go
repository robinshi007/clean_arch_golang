package graphql

import (
	"github.com/graphql-go/graphql"

	"clean_arch/infra"
	"clean_arch/infra/util"
)

// NewSchema -
func NewSchema(db infra.DB) graphql.Schema {
	var schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: NewRootQuery(db),
	})

	util.FailedIf(err)

	return schema
}
