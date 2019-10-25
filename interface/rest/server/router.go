package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"clean_arch/infra"
	"clean_arch/interface/rest/handler"
)

// NewRouter -
func NewRouter(db infra.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	uHanlder := handler.NewUserHandler(db)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", handler.NewUserRouter(uHanlder))
	})

	return r
}
