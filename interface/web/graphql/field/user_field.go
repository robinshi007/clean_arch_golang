package field

import (
	"context"
	"time"

	"github.com/graphql-go/graphql"

	"clean_arch/adapter/postgres"
	"clean_arch/adapter/presenter"
	"clean_arch/infra"
	"clean_arch/interface/web/graphql/types"
	"clean_arch/usecase"
)

// NewUserListField -
func NewUserListField(db infra.DB) *graphql.Field {
	userListField := &graphql.Field{
		Type:        graphql.NewList(types.UserType),
		Description: "List of users",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			repo := postgres.NewUserRepo(db)
			pre := presenter.NewUserPresenter()
			uc := usecase.NewUserUseCase(repo, pre, time.Second)
			return uc.GetAll(context.Background(), 5)
		},
	}
	return userListField
}
