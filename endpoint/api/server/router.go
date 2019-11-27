package server

import (
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"clean_arch/endpoint/api/handler"
	"clean_arch/infra"
)

// Hello -
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Clean Arch."))
}

// NewRouter -
func NewRouter(db infra.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	uHanlder := handler.NewUserHandler()
	aHanlder := handler.NewAccountHandler()

	r.Get("/", Hello)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", handler.NewUserRouter(uHanlder))
		rt.Mount("/accounts", handler.NewAccountRouter(aHanlder))
		rt.Mount("/play", gqlhandler.Playground("GraphQL Playground", "/graphql"))
		rt.Mount("/graphql", handler.GraphQLHandler())

	})

	return r
}
