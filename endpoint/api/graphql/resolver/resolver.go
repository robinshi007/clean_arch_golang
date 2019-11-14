package resolver

import (
	"context"
	"strconv"

	"clean_arch/domain/model"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api/graphql/gen"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver -
type Resolver struct {
	UserUC usecase.UserUsecase
}

// NewRootResolver -
func NewRootResolver(uuc usecase.UserUsecase) gen.Config {
	return gen.Config{
		Resolvers: &Resolver{
			UserUC: uuc,
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

// User -
func (r *Resolver) User() gen.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input in.NewUser) (*out.User, error) {
	userID, err := r.UserUC.Create(ctx, &input)
	if err != nil {
		return nil, err
	}
	user, _ := r.UserUC.GetByID(ctx, &in.FetchUser{ID: int64(userID)})
	return user, nil
}
func (r *mutationResolver) UpdateUser(ctx context.Context, input in.EditUserInput) (*out.User, error) {
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := r.UserUC.GetByID(ctx, &in.FetchUser{ID: id})
	if err != nil {
		return nil, err
	}
	euser := in.EditUser{
		User: &model.User{ID: user.ID, Name: user.Name},
		Name: input.Name,
	}
	return r.UserUC.Update(ctx, &euser)
}
func (r *mutationResolver) DeleteUser(ctx context.Context, input in.FetchUserInput) (*out.User, error) {
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := r.UserUC.GetByID(ctx, &in.FetchUser{ID: id})
	if err != nil {
		return nil, err
	}
	err = r.UserUC.Delete(ctx, &in.FetchUser{ID: id})
	if err != nil {
		return nil, err
	}
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*out.User, error) {
	return r.UserUC.GetAll(ctx, 10)
}
func (r *queryResolver) FetchUser(ctx context.Context, input in.FetchUserInput) (*out.User, error) {
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.UserUC.GetByID(ctx, &in.FetchUser{ID: id})
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *out.User) (string, error) {
	res := strconv.FormatInt(obj.ID, 10)
	return res, nil
}
