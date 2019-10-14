package repository

import (
	"clean_arch/domain/model"
)

// UserRepository -
type UserRepository interface {
	//Fetch(ctx context.Context, query string, args ...interface{}) ([]*model.User, error)
	Fetch(num int64) ([]*model.User, error)
	GetByID(id int64) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Create(u *model.User) (int64, error)
	Update(u *model.User) (*model.User, error)
	Delete(id int64) (bool, error)
}
