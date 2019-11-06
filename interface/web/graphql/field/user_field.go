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
	"clean_arch/usecase/input"
	"clean_arch/usecase/output"
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

// NewCreateUserField -
func NewCreateUserField(db infra.DB) *graphql.Field {

	CreateUserField := &graphql.Field{
		Type:        types.UserType,
		Description: "Create new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			userName, _ := params.Args["name"].(string)
			user := input.PostUser{Name: userName}

			if err := user.Validate(); err != nil {
				return nil, err
			}

			repo := postgres.NewUserRepo(db)
			pre := presenter.NewUserPresenter()
			uc := usecase.NewUserUseCase(repo, pre, time.Second)

			newID, err := uc.Create(context.Background(), &user)
			if err != nil {
				return nil, err
			}
			return output.User{ID: int64(newID), Name: userName}, nil
		},
	}
	return CreateUserField
}
