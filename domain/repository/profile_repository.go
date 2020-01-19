package repository

import (
	"context"

	"clean_arch/domain/model"
)

// ProfileRepository -
type ProfileRepository interface {
	FindAll(ctx context.Context, opt *ListOptions) ([]*model.UserProfile, error)
	FindByID(ctx context.Context, id int64) (*model.UserProfile, error)
	FindByEmail(ctx context.Context, email string) (*model.UserProfile, error)
}
