package presenter

import (
	"context"

	"clean_arch/domain/model"
	"clean_arch/domain/usecase/out"
)

// AccountPresenter -
type AccountPresenter interface {
	ViewError(ctx context.Context, err error) *out.Error
	ViewAccountID(ctx context.Context, user *model.UserAccount) out.ID
	ViewAccount(ctx context.Context, user *model.UserAccount) *out.Account
	ViewAccounts(ctx context.Context, users []*model.UserAccount) []*out.Account
}
