package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"clean_arch/endpoint/api/globals"

	"github.com/casbin/casbin"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
)

var (
	// ActionQuery -
	ActionQuery = "Query"
	// ActionMutation -
	ActionMutation = "Mutation"
)

// WithAuthorization -
func WithAuthorization(ef *casbin.Enforcer) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Read the content
			var bodyBytes []byte
			if r != nil {
				bodyBytes, _ = ioutil.ReadAll(r.Body)
			}
			// Restore the io.ReadCloser to its original state
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			var rb = struct {
				Query string `json:"query"`
			}{}
			json.Unmarshal(bodyBytes, &rb)

			// We have got subject from custom header
			// in real production subject could be in
			// something like jwt claim etc

			_, claims, _ := FromJWTContext(r.Context())

			var subject = claims.Email

			doc, _ := parser.Parse(parser.ParseParams{Source: rb.Query})
			for _, node := range doc.Definitions {
				switch d := node.(type) {
				case ast.TypeSystemDefinition:
					{
						var o string
						switch d.GetOperation() {
						case ast.OperationTypeQuery:
							o = ActionQuery
						case ast.OperationTypeMutation:
							o = ActionMutation
						default:
							continue
						}
						for _, s := range d.GetSelectionSet().Selections {
							switch f := s.(type) {
							case *ast.Field:
								if !ef.Enforce(subject, f.Name.Value, o) {
									fmt.Printf("ef: %v, %v, %v\n", subject, f.Name.Value, o)
									globals.Respond.GraphQLError(w, http.StatusOK, "the action is unauthorized", "authorization_checker")
									return
								}
							}
						}
					}
				}
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
