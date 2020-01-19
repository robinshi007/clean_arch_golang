package resolver

import (
	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api/graphql/gen"
	"clean_arch/infra/util"
	"context"
	"strconv"
)

// User -
func (r *Resolver) User() gen.UserResolver {
	return &userResolver{r}
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *out.User) (string, error) {
	res := strconv.FormatInt(obj.ID, 10)
	return res, nil
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *out.User) (string, error) {
	res := obj.CreatedAt.Format(util.TimeFormatStr)
	return res, nil
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *out.User) (string, error) {
	res := obj.UpdatedAt.Format(util.TimeFormatStr)
	return res, nil
}

// mutationResolver
func (r *mutationResolver) CreateUser(ctx context.Context, input in.NewUser) (*out.User, error) {
	userID, err := r.UserUC.Create(ctx, &input)
	if err != nil {
		return nil, err
	}
	user, _ := r.UserUC.FindByID(ctx, &in.FetchUser{ID: string(userID)})
	return user, nil
}
func (r *mutationResolver) UpdateUser(ctx context.Context, input in.EditUser) (*out.User, error) {
	return r.UserUC.Update(ctx, &input)
}
func (r *mutationResolver) DeleteUser(ctx context.Context, input in.FetchUser) (*out.User, error) {
	user, err := r.UserUC.FindByID(ctx, &input)
	if err != nil {
		return nil, err
	}
	err = r.UserUC.Delete(ctx, &input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// queryResolver -
func (r *queryResolver) Users(ctx context.Context) ([]*out.User, error) {
	return r.UserUC.FindAll(ctx, &in.FetchAllOptions{})
}
func (r *queryResolver) FetchUser(ctx context.Context, input in.FetchUser) (*out.User, error) {
	return r.UserUC.FindByID(ctx, &input)
}
