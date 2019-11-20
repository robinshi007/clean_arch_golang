package usecase

import (
	"context"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

// AccountUsecase -
type AccountUsecase interface {
	GetAll(ctx context.Context, num int64) ([]*out.Account, error)
	GetByID(ctx context.Context, u *in.FetchAccount) (*out.Account, error)
	GetByName(ctx context.Context, name string) (*out.Account, error)
	Create(ctx context.Context, u *in.NewAccount) (out.UserID, error)
	UpdatePassword(ctx context.Context, u *in.EditAccount) (*out.Account, error)
	Delete(ctx context.Context, u *in.FetchAccount) error
}
