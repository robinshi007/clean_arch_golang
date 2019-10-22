package presenter

import (
	"clean_arch/domain/model"
	"clean_arch/presenter/vm"
	"context"
)

// UserPresenter -
type UserPresenter interface {
	ViewError(ctx context.Context, err error) *vm.Error
	ViewUser(ctx context.Context, user *model.User) *vm.User
	ViewUsers(ctx context.Context, users []*model.User) []*vm.User
}
