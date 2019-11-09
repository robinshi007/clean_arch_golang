package repository

import (
	"context"

	"clean_arch/domain/model"
)

// UserListOptions -
type UserListOptions struct {
	Query string
	*LimitOffset
}

// UserRepository -
type UserRepository interface {
	// GetAll(ctx context.Context, query string, args ...interface{}) ([]*model.User, error)
	GetAll(ctx context.Context, opt *UserListOptions) ([]*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, u *model.User) (int64, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	DuplicatedByName(ctx context.Context, name string) error
}
