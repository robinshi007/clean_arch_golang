package presenter

import (
	"clean_arch/domain/model"
	"clean_arch/domain/presenter"
	"clean_arch/domain/usecase/out"
	"context"
)

// NewRedirectPresenter -
func NewRedirectPresenter() presenter.RedirectPresenter {
	return &redirectPresenter{}
}

type redirectPresenter struct {
}

// ViewError -
func (u redirectPresenter) ViewError(ctx context.Context, err error) *out.Error {
	return &out.Error{
		Code:    "500",
		Message: err.Error(),
	}
}

// ViewRedirect -
func (u redirectPresenter) ViewRedirect(ctx context.Context, redirect *model.Redirect) *out.Redirect {
	return &out.Redirect{
		ID:   redirect.ID,
		Code: redirect.Code,
		URL:  redirect.URL,
		CreatedBy: out.Profile{
			ID:        redirect.CreatedBy.UID,
			Name:      redirect.CreatedBy.Name,
			Email:     redirect.CreatedBy.Email,
			CreatedAt: redirect.CreatedBy.CreatedAt,
			UpdatedAt: redirect.CreatedBy.UpdatedAt,
		},
		CreatedAt: redirect.CreatedAt,
	}
}

// ViewRedirects -
func (u redirectPresenter) ViewRedirects(ctx context.Context, redirects []*model.Redirect) []*out.Redirect {
	res := make([]*out.Redirect, len(redirects))
	for i, redirect := range redirects {
		res[i] = &out.Redirect{
			ID:   redirect.ID,
			Code: redirect.Code,
			URL:  redirect.URL,
			CreatedBy: out.Profile{
				ID:        redirect.CreatedBy.UID,
				Name:      redirect.CreatedBy.Name,
				Email:     redirect.CreatedBy.Email,
				CreatedAt: redirect.CreatedBy.CreatedAt,
				UpdatedAt: redirect.CreatedBy.UpdatedAt,
			},
			CreatedAt: redirect.CreatedAt,
		}
	}
	return res
}
