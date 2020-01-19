package repository

import (
	"context"

	"clean_arch/domain/model"
)

// UserRepository -
type UserRepository interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, opt *ListOptions) ([]*model.User, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByName(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, u *model.User) (int64, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	ExistsByName(ctx context.Context, name string) error
}
