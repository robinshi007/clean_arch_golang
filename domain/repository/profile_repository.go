package repository

import (
	"context"

	"clean_arch/domain/model"
)

// ProfileListOptions -
type ProfileListOptions struct {
	Query string
	*LimitOffset
}

// ProfileRepository -
type ProfileRepository interface {
	GetAll(ctx context.Context, opt *ProfileListOptions) ([]*model.UserProfile, error)
	GetByID(ctx context.Context, id int64) (*model.UserProfile, error)
	GetByEmail(ctx context.Context, email string) (*model.UserProfile, error)
}
