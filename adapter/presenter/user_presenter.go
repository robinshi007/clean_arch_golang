package presenter

import (
	"clean_arch/domain/model"
	"clean_arch/domain/presenter"
	"clean_arch/domain/usecase/out"
	"context"
	"strconv"
)

// NewUserPresenter -
func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

type userPresenter struct {
}

// ViewError -
func (u userPresenter) ViewError(ctx context.Context, err error) *out.Error {
	return &out.Error{
		Code:    "500",
		Message: err.Error(),
	}
}

// ViewUserID -
func (u userPresenter) ViewUserID(ctx context.Context, user *model.User) out.ID {
	return out.ID(strconv.FormatInt(user.ID, 10))
}

// ViewUser -
func (u userPresenter) ViewUser(ctx context.Context, user *model.User) *out.User {
	return &out.User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ViewUsers -
func (u userPresenter) ViewUsers(ctx context.Context, users []*model.User) []*out.User {
	res := make([]*out.User, len(users))
	for i, user := range users {
		res[i] = &out.User{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return res
}
