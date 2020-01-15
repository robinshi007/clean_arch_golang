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
	urepo := postgres.NewUserRepo()
	upre := presenter.NewUserPresenter()
	uuc := usecase.NewUserUseCase(urepo, upre, time.Second)
	arepo := postgres.NewAccountRepo()
	apre := presenter.NewAccountPresenter()
	auc := usecase.NewAccountUseCase(arepo, apre, time.Second)
	rrepo := postgres.NewRedirectRepo()
	rpre := presenter.NewRedirectPresenter()
	ruc := usecase.NewRedirectUsecase(rrepo, rpre)
	return gqlhandler.GraphQL(gen.NewExecutableSchema(resolver.NewRootResolver(uuc, auc, ruc)))
}
