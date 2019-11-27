package resolver

import (
	"clean_arch/domain/usecase"
	"clean_arch/endpoint/api/graphql/gen"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver -
type Resolver struct {
	UserUC    usecase.UserUsecase
	AccountUC usecase.AccountUsecase
}

// NewRootResolver -
func NewRootResolver(
	uuc usecase.UserUsecase,
	auc usecase.AccountUsecase,
) gen.Config {
	return gen.Config{
		Resolvers: &Resolver{
			UserUC:    uuc,
			AccountUC: auc,
		},
	}

}

// Mutation -
func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}

// Query -
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
