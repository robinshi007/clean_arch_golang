package presenter

import (
	"context"

	"clean_arch/domain/model"
	"clean_arch/domain/usecase/out"
)

// RedirectPresenter -
type RedirectPresenter interface {
	ViewError(ctx context.Context, err error) *out.Error
	ViewRedirect(ctx context.Context, redirect *model.Redirect) *out.Redirect
	ViewRedirects(ctx context.Context, redirects []*model.Redirect) []*out.Redirect
}
