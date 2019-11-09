package server

import (
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	gql "clean_arch/endpoint/api/graphql"
	"clean_arch/endpoint/api/handler"
	"clean_arch/infra"
	"clean_arch/infra/util"
)

// NewRouter -
func NewRouter(db infra.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	uHanlder := handler.NewUserHandler(db)
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	util.FailedIf(err)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", handler.NewUserRouter(uHanlder))

		rt.Mount("/graphql", gql.NewGraphqlHandler(db))
		rt.Mount("/graphiql", graphiqlHandler)
	})

	return r
}
