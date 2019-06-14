package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/robinshi007/goweb/db"
)

func NewRouter(conn *db.Db) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	uHanlder := NewUserHandler(conn)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", NewUserRouter(uHanlder))
	})

	return r
}
