package usecase

import (
	"context"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

// UserUsecase -
type UserUsecase interface {
	GetAll(ctx context.Context, num int64) ([]*out.User, error)
	GetByID(ctx context.Context, u *in.FetchUser) (*out.User, error)
	GetByName(ctx context.Context, name string) (*out.User, error)
	Create(ctx context.Context, u *in.NewUser) (out.UserID, error)
	Update(ctx context.Context, u *in.EditUser) (*out.User, error)
	Delete(ctx context.Context, u *in.FetchUser) error
}
