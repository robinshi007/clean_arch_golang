package server

import (
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"

	"clean_arch/endpoint/api/handler"
	mw "clean_arch/endpoint/api/middleware"
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
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Logger)

	//uHanlder := handler.NewUserHandler()
	aHanlder := handler.NewAccountHandler()
	authHanlder := handler.NewAuthHandler()

	r.Get("/", Hello)
	r.Route("/", func(rt chi.Router) {
		rt.Use(mw.JWTMiddleware())
		//rt.Mount("/user", handler.NewUserRouter(uHanlder))
		rt.Mount("/auth", handler.NewAuthRouter(authHanlder))
		rt.Mount("/accounts", handler.NewAccountRouter(aHanlder))
		rt.Mount("/play", gqlhandler.Playground("GraphQL Playground", "/graphql"))
		rt.Mount("/graphql", handler.GraphQLHandler())

	})

	return r
}
