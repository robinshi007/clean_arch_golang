package repository

import (
	"github.com/robinshi007/goweb/domain/model"
)

type UserRepository interface {
	//Fetch(ctx context.Context, query string, args ...interface{}) ([]*model.User, error)
	Fetch(num int64) ([]*model.User, error)
	GetById(id int64) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Create(u *model.User) (int64, error)
	Update(u *model.User) (*model.User, error)
	Delete(id int64) (bool, error)
}
