package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/robinshi007/goweb/db"
	uh "github.com/robinshi007/goweb/handler"
)

func main() {
	conn, err := db.NewDb("localhost", "5432", "postgres", "postgres", "test")
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	uHanlder := uh.NewUserHandler(conn)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", userRouter(uHanlder))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}

func userRouter(uHandler *uh.User) http.Handler {
	r := chi.NewRouter()
	r.Get("/", uHandler.Fetch)
	r.Get("/{id:[0-9]+}", uHandler.GetById)
	r.Post("/", uHandler.Create)
	r.Put("/{id:[0-9]+}", uHandler.Update)
	r.Delete("/{id:[0-9]+}", uHandler.Delete)
	return r
}
