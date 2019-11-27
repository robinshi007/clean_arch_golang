package usecase

import (
	"context"
	"strconv"
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

func (u *userUsecase) GetByID(c context.Context, input *in.FetchUser) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}
	id, err := in.ToID(input.ID)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUser(ctx, user), nil
}

func (u *userUsecase) GetByName(c context.Context, input *in.FetchUserByName) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}

	user, err := u.repo.GetByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUser(ctx, user), nil
}

func (u *userUsecase) Create(c context.Context, input *in.NewUser) (out.ID, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	if err := in.Validate(input); err != nil {
		return out.ID("-1"), model.ErrEntityBadInput
	}

	res, err := u.repo.Create(ctx, &model.User{
		Name: input.Name,
	})
	id := strconv.FormatInt(res, 10)
	return out.ID(id), err
}

func (u *userUsecase) Update(c context.Context, input *in.EditUser) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	if err := in.Validate(input); err != nil {
		return nil, model.ErrEntityBadInput
	}
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, model.ErrEntityBadInput
	}
	user, err := u.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, model.ErrEntityNotFound
	}

	// check is dirty
	if user.Name == input.Name {
		return nil, model.ErrEntityNotChanged
	}
	usr, err := u.repo.Update(ctx, &model.User{
		ID:   user.ID,
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return u.pre.ViewUser(ctx, usr), nil
}

func (u *userUsecase) Delete(c context.Context, input *in.FetchUser) error {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err := in.Validate(input)
	if err != nil {
		return model.ErrEntityBadInput
	}

	id, err := in.ToID(input.ID)
	if err != nil {
		return err
	}
	return u.repo.Delete(ctx, id)
}

func (u *userUsecase) RegisterUser(c context.Context, name string) (out.ID, error) {
	if err := u.repo.DuplicatedByName(c, name); err != nil {
		return out.ID("-1"), err
	}
	user := model.NewUser(name)
	id, err := u.repo.Create(c, user)
	if err != nil {
		return out.ID("-1"), err
	}
	return out.ID(strconv.FormatInt(id, 10)), nil
}
