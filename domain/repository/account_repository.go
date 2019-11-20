package repository

import (
	"context"

	"clean_arch/domain/model"
)

// AccountListOptions -
type AccountListOptions struct {
	Query string
	*LimitOffset
}

// AccountRepository -
type AccountRepository interface {
	GetAll(ctx context.Context, opt *AccountListOptions) ([]*model.UserAccount, error)
	GetByID(ctx context.Context, id int64) (*model.UserAccount, error)
	GetByEmail(ctx context.Context, email string) (*model.UserAccount, error)
	GetByName(ctx context.Context, name string) (*model.UserAccount, error)
	Create(ctx context.Context, u *model.UserAccount) (int64, error)
	Update(ctx context.Context, u *model.UserAccount) (*model.UserAccount, error)
	Delete(ctx context.Context, id int64) error
	DuplicatedByEmail(ctx context.Context, email string) error
	DuplicatedByName(ctx context.Context, name string) error
}
