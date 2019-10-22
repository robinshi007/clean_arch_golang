package repository

import (
	"clean_arch/domain/model"
	"context"
)

// UserRepository -
type UserRepository interface {
	//Fetch(ctx context.Context, query string, args ...interface{}) ([]*model.User, error)
	Fetch(ctx context.Context, num int64) ([]*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, u *model.User) (int64, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	DuplicatedByName(ctx context.Context, name string) error
}
