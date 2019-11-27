package presenter

import (
	"context"

	"clean_arch/domain/model"
	"clean_arch/domain/usecase/out"
)

// UserPresenter -
type UserPresenter interface {
	ViewError(ctx context.Context, err error) *out.Error
	ViewUserID(ctx context.Context, user *model.User) out.ID
	ViewUser(ctx context.Context, user *model.User) *out.User
	ViewUsers(ctx context.Context, users []*model.User) []*out.User
}
