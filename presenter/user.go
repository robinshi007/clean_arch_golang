package presenter

import (
	"clean_arch/domain/model"
	"clean_arch/presenter/vm"
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
func (u userPresenter) ViewError(ctx context.Context, err error) *vm.Error {
	return &vm.Error{
		Code:    500,
		Message: err.Error(),
	}
}

// ViewUser -
func (u userPresenter) ViewUser(ctx context.Context, user *model.User) *vm.User {
	return &vm.User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}
}

// ViewUsers -
func (u userPresenter) ViewUsers(ctx context.Context, users []*model.User) []*vm.User {
	res := make([]*vm.User, len(users))
	for i, user := range users {
		res[i] = &vm.User{
			ID:   user.GetID(),
			Name: user.GetName(),
		}
	}
	return res
}
