package repository

import (
	"context"

	"clean_arch/domain/model"
)

// AccountRepository -
type AccountRepository interface {
	FindAll(ctx context.Context, opt *ListOptions) ([]*model.UserAccount, error)
	FindByID(ctx context.Context, id int64) (*model.UserAccount, error)
	FindByEmail(ctx context.Context, email string) (*model.UserAccount, error)
	FindByName(ctx context.Context, name string) (*model.UserAccount, error)
	Create(ctx context.Context, u *model.UserAccount) (int64, error)
	Update(ctx context.Context, u *model.UserAccount) (*model.UserAccount, error)
	UpdatePassword(ctx context.Context, u *model.UserAccount) (*model.UserAccount, error)
	Delete(ctx context.Context, id int64) error
	ExistsByEmail(ctx context.Context, email string) error
	ExistsByName(ctx context.Context, name string) error
}
