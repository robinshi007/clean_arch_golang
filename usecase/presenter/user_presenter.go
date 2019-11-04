package presenter

import (
	"context"

	"clean_arch/domain/model"
	out "clean_arch/usecase/output"
)

// UserPresenter -
type UserPresenter interface {
	ViewError(ctx context.Context, err error) *out.Error
	ViewUserID(ctx context.Context, user *model.User) out.UserID
	ViewUser(ctx context.Context, user *model.User) *out.User
	ViewUsers(ctx context.Context, users []*model.User) []*out.User
}
