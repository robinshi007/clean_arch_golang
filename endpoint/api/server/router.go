package server

import (
	"net/http"
	"time"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/casbin/casbin"
	"github.com/go-chi/chi"
	chiMW "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"clean_arch/endpoint/api/handler"
	mw "clean_arch/endpoint/api/middleware"
	"clean_arch/infra"
	"clean_arch/registry"
)

// NewRouter -
func NewRouter(db infra.DB) http.Handler {

	// middlewares
	// cors
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// casbin
	e := casbin.NewEnforcer("./config/casbin_model.conf")
	// Having public access for feching schema
	e.AddPolicy("public", "__schema", mw.ActionQuery)
	// Alice just has access to foo query
	e.AddPolicy("admin@test.com", "accounts", mw.ActionQuery)
	e.AddPolicy("admin@test.com", "fetchAccount", mw.ActionQuery)
	e.AddPolicy("admin@test.com", "createAccount", mw.ActionMutation)
	e.AddPolicy("admin@test.com", "updateAccount", mw.ActionMutation)
	e.AddPolicy("admin@test.com", "deleteAccount", mw.ActionMutation)

	// admin have access to public resources
	e.AddGroupingPolicy("admin@test.com", "public")

	// handlers
	//uHanlder := handler.NewUserHandler()
	eHandler := handler.NewErrorHandler()
	//aHandler := handler.NewAccountHandler()
	auHandler := handler.NewAuthHandler()

	r := chi.NewRouter()
	//	r.Use(chiMW.RequestID)
	//	r.Use(chiMW.RealIP)
	r.Use(chiMW.Logger)
	//r.Use(chiMW.Recoverer)
	r.Use(eHandler.Recoverer)
	r.Use(chiMW.Timeout(10 * time.Second))
	r.Use(cors.Handler)

	// for test only
	r.Get("/", handler.HelloHanlder)
	r.Get("/panic", handler.PanicHanlder)

	// for graphql playground, dev only, will removed in prod env
	if registry.Cfg.Mode == "dev" {
		r.Mount("/play", gqlhandler.Playground("GraphQL Playground", "/api/v1/graphql"))
	}

	// for api use
	r.Route("/api/v1", func(rt chi.Router) {
		rt.Use(mw.JWTVerify())
		// for login
		rt.Post("/auth/login", auHandler.Login)
		// for refresh token
		rt.Mount("/auth", mw.New(mw.JWTAuthenticator).Then(handler.NewAuthRouter(auHandler)))
		//rt.Mount("/user", handler.NewUserRouter(uHanlder))
		//rt.Mount("/accounts", mw.New(mw.JWTAuthenticator).Then(handler.NewAccountRouter(aHandler)))
		rt.Mount("/graphql", mw.New(mw.JWTAuthenticator, mw.WithAuthorization(e)).Then(handler.GraphQLHandler()))
		//rt.Mount("/graphql", handler.GraphQLHandler())

	})

	// Error Handler
	r.NotFound(eHandler.RouteNotFound)
	r.MethodNotAllowed(eHandler.MethodNotAllowed)
	return r
}
