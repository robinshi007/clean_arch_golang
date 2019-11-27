package resolver

import (
	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api/graphql/gen"
	"context"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Account() gen.AccountResolver {
	return &accountResolver{r}
}
func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() gen.UserResolver {
	return &userResolver{r}
}

type accountResolver struct{ *Resolver }

func (r *accountResolver) ID(ctx context.Context, obj *out.Account) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input in.NewUser) (*out.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, input in.EditUser) (*out.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, input in.FetchUser) (*out.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateAccount(ctx context.Context, input in.NewAccount) (*out.Account, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAccount(ctx context.Context, input in.EditAccount) (*out.Account, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAccount(ctx context.Context, input in.FetchAccount) (*out.Account, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*out.User, error) {
	panic("not implemented")
}
func (r *queryResolver) FetchUser(ctx context.Context, input in.FetchUser) (*out.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Accounts(ctx context.Context) ([]*out.Account, error) {
	panic("not implemented")
}
func (r *queryResolver) FetchAccount(ctx context.Context, input in.FetchAccount) (*out.Account, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *out.User) (string, error) {
	panic("not implemented")
}
