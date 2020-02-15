package usecase

import (
	"context"
	"errors"
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
	"clean_arch/infra/util"
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
func (r *redirectUsecase) FindAll(ctx context.Context, input *in.FetchAllOptions) ([]*out.Redirect, error) {
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

	redirects, err := r.repo.FindAll(ctx, &repository.ListOptions{
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

func (r *redirectUsecase) FindByURL(ctx context.Context, input *in.FetchRedirectByURL) (*out.Redirect, error) {
	if err := in.Validate(input); err != nil {
		return nil, fmt.Errorf("redirectUsecase.FindByURL: %w", model.ErrEntityBadInput)
	}
	redirect, err := r.repo.FindByURL(ctx, input.URL)
	if err != nil {
		return nil, err
	}
	return r.pre.ViewRedirect(ctx, redirect), nil
}

func (r *redirectUsecase) FindOrCreate(ctx context.Context, input *in.FetchOrCreateRedirect) (*out.Redirect, error) {
	if err := in.Validate(input); err != nil {
		return nil, fmt.Errorf("redirectUsecase.FindByURL: %w", model.ErrEntityBadInput)
	}
	redirect, err := r.repo.FindByURL(ctx, input.URL)
	if err != nil {
		if errors.Is(err, model.ErrEntityNotFound) {
			newRedirectID, err := r.Create(ctx, &in.NewRedirect{
				URL: input.URL,
				CID: input.CID,
			})
			if err != nil {
				return nil, err
			}
			newID, _ := in.ToID(string(newRedirectID))
			newRedirect, err := r.repo.FindByID(ctx, newID)
			if err != nil {
				return nil, err
			}
			return r.pre.ViewRedirect(ctx, newRedirect), nil
		}
		return nil, err
	}
	return r.pre.ViewRedirect(ctx, redirect), nil
}

func (r *redirectUsecase) Create(ctx context.Context, input *in.NewRedirect) (out.ID, error) {
	if err := in.Validate(input); err != nil {
		return out.ID("-1"), fmt.Errorf("redirectUsecase.Create: %w", model.ErrEntityBadInput)
	}
	cid, _ := util.String2Int64(input.CID)
	res, err := r.repo.Create(ctx, &model.Redirect{
		Code:      shortid.MustGenerate(),
		URL:       input.URL,
		CreatedBy: model.UserProfile{UID: cid},
		CreatedAt: time.Now(),
	})

	id := strconv.FormatInt(res, 10)
	return out.ID(id), err
}
