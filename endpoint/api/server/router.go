package server

import (
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"clean_arch/endpoint/api/handler"
	"clean_arch/infra"
)

// NewRouter -
func NewRouter(db infra.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	uHanlder := handler.NewUserHandler()

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", handler.NewUserRouter(uHanlder))
		rt.Mount("/play", gqlhandler.Playground("GraphQL Playground", "/graphql"))
		rt.Mount("/graphql", handler.GraphQLHandler())

	})

	return r
}
