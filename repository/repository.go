package repository

import (
	"context"

	"github.com/robinshi007/goweb/model"
)

type UserRepo interface {
	//Fetch(ctx context.Context, query string, args ...interface{}) ([]*model.User, error)
	Fetch(ctx context.Context, num int64) ([]*model.User, error)
	GetById(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, u *model.User) (int64, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
