package usecase

import (
	"context"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

// AccountUsecase -
type AccountUsecase interface {
	GetAll(ctx context.Context, num int64) ([]*out.Account, error)
	GetByID(ctx context.Context, input *in.FetchAccount) (*out.Account, error)
	GetByEmail(ctx context.Context, input *in.FetchAccountByEmail) (*out.Account, error)
	Create(ctx context.Context, input *in.NewAccount) (out.ID, error)
	UpdatePassword(ctx context.Context, input *in.EditAccount) (*out.Account, error)
	Delete(ctx context.Context, input *in.FetchAccount) error
	Login(ctx context.Context, input *in.LoginAccountByEmail) (bool, error)
}
