package server

import (
	"net/http"
	"time"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/endpoint/api/graphql/gen"
	"clean_arch/endpoint/api/graphql/resolver"
	"clean_arch/endpoint/api/handler"
	"clean_arch/infra"
	"clean_arch/usecase"
)

// NewRouter -
func NewRouter(db infra.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	uHanlder := handler.NewUserHandler()

	repo := postgres.NewUserRepo()
	pre := presenter.NewUserPresenter()
	uuc := usecase.NewUserUseCase(repo, pre, time.Second)
	//	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	//	util.FailedIf(err)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", handler.NewUserRouter(uHanlder))
		rt.Mount("/play", gqlhandler.Playground("GraphQL Playground", "/graphql"))
		rt.Mount("/graphql", gqlhandler.GraphQL(gen.NewExecutableSchema(resolver.NewRootResolver(uuc))))

		//		rt.Mount("/graphql", gql.NewGraphqlHandler())
		//		rt.Mount("/graphiql", graphiqlHandler)
	})

	return r
}
