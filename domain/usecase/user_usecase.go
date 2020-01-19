package usecase

import (
	"context"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

// UserUsecase -
type UserUsecase interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, input *in.FetchAllOptions) ([]*out.User, error)
	FindByID(ctx context.Context, input *in.FetchUser) (*out.User, error)
	FindByName(ctx context.Context, input *in.FetchUserByName) (*out.User, error)
	Create(ctx context.Context, input *in.NewUser) (out.ID, error)
	Update(ctx context.Context, input *in.EditUser) (*out.User, error)
	Delete(ctx context.Context, input *in.FetchUser) error
}
