package presenter

import (
	"clean_arch/domain/model"
	out "clean_arch/usecase/output"
	"clean_arch/usecase/presenter"
	"context"
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

// ViewUser -
func (u userPresenter) ViewUser(ctx context.Context, user *model.User) *out.User {
	return &out.User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}
}

// ViewUsers -
func (u userPresenter) ViewUsers(ctx context.Context, users []*model.User) []*out.User {
	res := make([]*out.User, len(users))
	for i, user := range users {
		res[i] = &out.User{
			ID:   user.GetID(),
			Name: user.GetName(),
		}
	}
	return res
}
