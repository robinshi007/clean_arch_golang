package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"

	"clean_arch/infra"
	"clean_arch/infra/util"
)

type reqBody struct {
	Query string `json:"query"`
}

// NewGraphqlHandler -
func NewGraphqlHandler(db infra.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		//body, err := ioutil.ReadAll(r.Body)

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		util.FailedIf(err)

		result := executeQuery(fmt.Sprintf("%s", rBody.Query), NewSchema(db))
		json.NewEncoder(w).Encode(result)
	})
}
func executeQuery(query string, schema graphql.Schema) *graphql.Result {

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected graphql errors: %v", result.Errors)
	}
	return result
}
