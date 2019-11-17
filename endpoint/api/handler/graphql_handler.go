package handler

import (
	"net/http"
	"time"

	gqlhandler "github.com/99designs/gqlgen/handler"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/endpoint/api/graphql/gen"
	"clean_arch/endpoint/api/graphql/resolver"
	"clean_arch/usecase"
)

// GraphQLHandler -
func GraphQLHandler() http.Handler {
	repo := postgres.NewUserRepo()
	pre := presenter.NewUserPresenter()
	uuc := usecase.NewUserUseCase(repo, pre, time.Second)
	return gqlhandler.GraphQL(gen.NewExecutableSchema(resolver.NewRootResolver(uuc)))
}
