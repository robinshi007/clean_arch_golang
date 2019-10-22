package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"clean_arch/infra/database"
)

// NewRouter -
func NewRouter(dbm database.DBM) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	uHanlder := NewUserHandler(dbm)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", NewUserRouter(uHanlder))
	})

	return r
}
