package server

import (
	"net/http"
	"time"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"clean_arch/endpoint/api/handler"
	"clean_arch/infra"
	"clean_arch/registry"
)

// NewRouter -
func NewRouter(db infra.DB) http.Handler {

	// middlewares
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// handlers
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
	r.Use(cors.Handler)

	// for test only
	r.Get("/", handler.HelloHanlder)
	r.Get("/panic", handler.PanicHanlder)

	// for graphql playground, dev only, will removed in prod env
	if registry.Cfg.Mode == "dev" {
		r.Mount("/play", gqlhandler.Playground("GraphQL Playground", "/api/v1/graphql"))
	}

	//	r.Post("/login", auHandler.Login)
	//	r.Route("/", func(rt chi.Router) {
	//		rt.Use(mw.JWTVerify())
	//		rt.Mount("/auth", auHandler.JWTAuthenticator(handler.NewAuthRouter(auHandler)))
	//	})

	// for api use
	r.Route("/api/v1", func(rt chi.Router) {
		rt.Use(auHandler.JWTVerify())
		// for login
		rt.Post("/auth/login", auHandler.Login)
		// for refresh token
		rt.Mount("/auth", auHandler.JWTAuthenticator(handler.NewAuthRouter(auHandler)))
		//rt.Mount("/user", handler.NewUserRouter(uHanlder))
		rt.Mount("/accounts", auHandler.JWTAuthenticator(handler.NewAccountRouter(aHandler)))
		rt.Mount("/graphql", auHandler.JWTAuthenticator(handler.GraphQLHandler()))
		//rt.Mount("/graphql", handler.GraphQLHandler())

	})

	// Error Handler
	r.NotFound(eHandler.RouteNotFound)
	r.MethodNotAllowed(eHandler.MethodNotAllowed)
	return r
}
