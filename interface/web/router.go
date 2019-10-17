package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"clean_arch/infra/database"
)

// NewRouter -
func NewRouter(db database.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	uHanlder := NewUserHandler(db)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", NewUserRouter(uHanlder))
	})

	return r
}
