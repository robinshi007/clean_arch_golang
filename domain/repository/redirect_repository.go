package repository

import (
	"context"

	"clean_arch/domain/model"
)

// RedirectListOptions -
type RedirectListOptions struct {
	Query       string
	LimitOffset *LimitOffset
}

// RedirectRepository -
type RedirectRepository interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, opt *RedirectListOptions) ([]*model.Redirect, error)
	FindByID(ctx context.Context, id int64) (*model.Redirect, error)
	FindByCode(ctx context.Context, code string) (*model.Redirect, error)
	Save(ctx context.Context, r *model.Redirect) (int64, error)
	Delete(ctx context.Context, id int64) error
}

// CRUD interface by spring-data JpaRepository
// Count() int64, error
// Exists(id) bool, error
// ExistsByName() <entity>, error
// FindAll(offset, page_size, sort) [entity], error
// Find(id) <entity>, error
// FindByName(name) <entity>, error
// Save(<entity>) <entity>, error
// Delete(id) <entity>, error
// DeleteInBatch(ids) [entity], error
