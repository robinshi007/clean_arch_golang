package usecase

import (
	"context"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

// UserUsecase -
type UserUsecase interface {
	GetAll(ctx context.Context, num int64) ([]*out.User, error)
	GetByID(ctx context.Context, input *in.FetchUser) (*out.User, error)
	GetByName(ctx context.Context, input *in.FetchUserByName) (*out.User, error)
	Create(ctx context.Context, input *in.NewUser) (out.ID, error)
	Update(ctx context.Context, input *in.EditUser) (*out.User, error)
	Delete(ctx context.Context, input *in.FetchUser) error
}
