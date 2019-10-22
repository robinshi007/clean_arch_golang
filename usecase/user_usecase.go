package usecase

import (
	"context"
	"time"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/presenter/vm"
	"clean_arch/usecase/in"
	"clean_arch/usecase/presenter"
)

// UserUsecase -
type UserUsecase interface {
	Fetch(ctx context.Context, num int64) ([]*vm.User, error)
	GetByID(ctx context.Context, id int64) (*vm.User, error)
	GetByName(ctx context.Context, name string) (*vm.User, error)
	Create(ctx context.Context, u *in.PostUser) (int64, error)
	Update(ctx context.Context, u *in.PutUser) (*vm.User, error)
	Delete(ctx context.Context, id int64) error
}

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
) UserUsecase {
	return &userUsecase{
		repo:       repo,
		pre:        pre,
		ctxTimeout: timeout,
	}
}

func (u *userUsecase) Fetch(c context.Context, num int64) ([]*vm.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	users, err := u.repo.Fetch(ctx, num)
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUsers(ctx, users), nil
}

func (u *userUsecase) GetByID(c context.Context, id int64) (*vm.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUser(ctx, user), nil
}

func (u *userUsecase) GetByName(c context.Context, name string) (*vm.User, error) {
	return nil, nil
}

func (u *userUsecase) Create(c context.Context, user *in.PostUser) (int64, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	return u.repo.Create(ctx, &model.User{
		Name: user.Name,
	})
}

func (u *userUsecase) Update(c context.Context, user *in.PutUser) (*vm.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	usr, err := u.repo.Update(ctx, &model.User{
		ID:   user.ID,
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

func (u *userUsecase) RegisterUser(c context.Context, name string) (int64, error) {
	if err := u.repo.DuplicatedByName(c, name); err != nil {
		return -1, err
	}
	user := model.NewUser(name)
	if id, err := u.repo.Create(c, user); err != nil {
		return id, err
	}
	return -1, nil
}
