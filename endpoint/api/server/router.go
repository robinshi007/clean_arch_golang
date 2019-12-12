package server

import (
	"net/http"
	"time"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"

	"clean_arch/endpoint/api/handler"
	mw "clean_arch/endpoint/api/middleware"
	"clean_arch/infra"
	"clean_arch/registry"
)

// NewRouter -
func NewRouter(db infra.DB) http.Handler {
	//uHanlder := handler.NewUserHandler()
	eHandler := handler.NewErrorHandler()
	aHandler := handler.NewAccountHandler()
	auHandler := handler.NewAuthHandler()

	r := chi.NewRouter()
	//	r.Use(chiMiddleware.RequestID)
	//	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	//r.Use(chiMiddleware.Recoverer)
	r.Use(eHandler.Recoverer)
	r.Use(chiMiddleware.Timeout(10 * time.Second))

	// for test only
	r.Get("/", handler.HelloHanlder)
	r.Get("/panic", handler.PanicHanlder)

	// for graphql playground, dev only, will removed in prod env
	if registry.Cfg.Mode == "dev" {
		r.Mount("/play", gqlhandler.Playground("GraphQL Playground", "/api/v1/graphql"))
	}

	// for login
	r.Post("/login", auHandler.Login)
	r.Route("/", func(rt chi.Router) {
		rt.Use(mw.JWTVerify())
		rt.Mount("/auth", auHandler.JWTAuthenticator(handler.NewAuthRouter(auHandler)))
	})

	// for api use
	r.Route("/api/v1", func(rt chi.Router) {
		rt.Use(mw.JWTVerify())
		//rt.Mount("/user", handler.NewUserRouter(uHanlder))
		rt.Mount("/accounts", auHandler.JWTAuthenticator(handler.NewAccountRouter(aHandler)))
		rt.Mount("/graphql", auHandler.JWTAuthenticator(handler.GraphQLHandler()))

	})

	// Error Handler
	r.NotFound(eHandler.RouteNotFound)
	r.MethodNotAllowed(eHandler.MethodNotAllowed)
	return r
}
