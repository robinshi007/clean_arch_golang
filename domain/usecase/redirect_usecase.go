package usecase

import (
	"context"

	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

// RedirectUsecase -
type RedirectUsecase interface {
	Count(ctx context.Context) (int64, error)
	FindAll(ctx context.Context, input *in.FetchAllOptions) ([]*out.Redirect, error)
	FindByID(ctx context.Context, input *in.FetchRedirect) (*out.Redirect, error)
	FindByCode(ctx context.Context, input *in.FetchRedirectByCode) (*out.Redirect, error)
	Create(ctx context.Context, input *in.NewRedirect) (out.ID, error)
	FindOrCreate(ctx context.Context, input *in.FetchOrCreateRedirect) (*out.Redirect, error)
}
