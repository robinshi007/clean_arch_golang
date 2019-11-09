package usecase

import (
	"context"
	"errors"
	"time"

	"clean_arch/domain/model"
	"clean_arch/domain/presenter"
	"clean_arch/domain/repository"
	"clean_arch/domain/usecase"
	"clean_arch/domain/usecase/in"
	"clean_arch/domain/usecase/out"
)

type userUsecase struct {
	repo       repository.UserRepository
	pre        presenter.UserPresenter
	ctxTimeout time.Duration
}

// NewUserUseCase -
func NewUserUseCase(
	repo repository.UserRepository,
	pre presenter.UserPresenter,
	timeout time.Duration,
) usecase.UserUsecase {
	return &userUsecase{
		repo:       repo,
		pre:        pre,
		ctxTimeout: timeout,
	}
}

func (u *userUsecase) GetAll(c context.Context, num int64) ([]*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	users, err := u.repo.GetAll(ctx, &repository.UserListOptions{})
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUsers(ctx, users), nil
}

func (u *userUsecase) GetByID(c context.Context, id int64) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUser(ctx, user), nil
}

func (u *userUsecase) GetByName(c context.Context, name string) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	user, err := u.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUser(ctx, user), nil
}

func (u *userUsecase) Create(c context.Context, user *in.PostUser) (out.UserID, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	res, err := u.repo.Create(ctx, &model.User{
		Name: user.Name,
	})
	return out.UserID(res), err
}

func (u *userUsecase) Update(c context.Context, user *in.PutUser) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	name := user.User.Name
	// check is dirty
	if name == user.Name {
		return nil, errors.New("item is not changed")
	}
	usr, err := u.repo.Update(ctx, &model.User{
		ID:   user.User.ID,
		Name: user.Name,
	})
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUser(ctx, usr), nil
}

func (u *userUsecase) Delete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	return u.repo.Delete(ctx, id)
}

func (u *userUsecase) RegisterUser(c context.Context, name string) (out.UserID, error) {
	if err := u.repo.DuplicatedByName(c, name); err != nil {
		return -1, err
	}
	user := model.NewUser(name)
	if id, err := u.repo.Create(c, user); err != nil {
		return out.UserID(id), err
	}
	return -1, nil
}
