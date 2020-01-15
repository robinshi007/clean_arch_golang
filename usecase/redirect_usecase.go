package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/teris-io/shortid"

	"clean_arch/domain/model"
	"clean_arch/domain/presenter"
	"clean_arch/domain/repository"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

type redirectUsecase struct {
	repo repository.RedirectRepository
	pre  presenter.RedirectPresenter
}

// NewRedirectUsecase -
func NewRedirectUsecase(
	repo repository.RedirectRepository,
	pre presenter.RedirectPresenter,
) usecase.RedirectUsecase {
	return &redirectUsecase{
		repo,
		pre,
	}
}

// Count -
func (r *redirectUsecase) Count(ctx context.Context) (int64, error) {
	return r.repo.Count(ctx)
}

// FindAll -
func (r *redirectUsecase) FindAll(ctx context.Context, input *in.FetchRedirects) ([]*out.Redirect, error) {
	if input.Offset == "" {
		input.Offset = "0"
	}
	if input.Limit == "" {
		input.Limit = "10"
	}
	if err := in.Validate(input); err != nil {
		return nil, fmt.Errorf("redirectUsecase.FindAll: %w", model.ErrEntityBadInput)
	}
	offset, _ := strconv.Atoi(input.Offset)
	limit, _ := strconv.Atoi(input.Limit)

	redirects, err := r.repo.FindAll(ctx, &repository.RedirectListOptions{
		Query: "",
		LimitOffset: &repository.LimitOffset{
			Limit:  limit,
			Offset: offset,
		},
	})
	if err != nil {
		return nil, err
	}
	return r.pre.ViewRedirects(ctx, redirects), nil
}
func (r *redirectUsecase) FindByID(ctx context.Context, input *in.FetchRedirect) (*out.Redirect, error) {
	if err := in.Validate(input); err != nil {
		return nil, fmt.Errorf("redirectUsecase.FindByID: %w", model.ErrEntityBadInput)
	}
	id, err := in.ToID(input.ID)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	redirect, err := r.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.pre.ViewRedirect(ctx, redirect), nil
}

func (r *redirectUsecase) FindByCode(ctx context.Context, input *in.FetchRedirectByCode) (*out.Redirect, error) {
	if err := in.Validate(input); err != nil {
		return nil, fmt.Errorf("redirectUsecase.FindByCode: %w", model.ErrEntityBadInput)
	}
	redirect, err := r.repo.FindByCode(ctx, input.Code)
	if err != nil {
		return nil, err
	}
	return r.pre.ViewRedirect(ctx, redirect), nil
}

func (r *redirectUsecase) Save(ctx context.Context, input *in.NewRedirect) (out.ID, error) {
	if err := in.Validate(input); err != nil {
		return out.ID("-1"), fmt.Errorf("redirectUsecase.Save: %w", model.ErrEntityBadInput)
	}

	res, err := r.repo.Save(ctx, &model.Redirect{
		Code:      shortid.MustGenerate(),
		URL:       input.URL,
		CreatedAt: time.Now(),
	})

	id := strconv.FormatInt(res, 10)
	return out.ID(id), err
}
