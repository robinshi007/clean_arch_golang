package usecase

import (
	"context"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

// AccountUsecase -
type AccountUsecase interface {
	FindAll(ctx context.Context, input *in.FetchAllOptions) ([]*out.Account, error)
	FindByID(ctx context.Context, input *in.FetchAccount) (*out.Account, error)
	FindByEmail(ctx context.Context, input *in.FetchAccountByEmail) (*out.Account, error)
	Create(ctx context.Context, input *in.NewAccount) (out.ID, error)
	Update(ctx context.Context, input *in.EditAccount) (*out.Account, error)
	UpdatePassword(ctx context.Context, input *in.EditAccountPassword) (*out.Account, error)
	Delete(ctx context.Context, input *in.FetchAccount) error
	Login(ctx context.Context, input *in.LoginAccountByEmail) (bool, string, error)
}
