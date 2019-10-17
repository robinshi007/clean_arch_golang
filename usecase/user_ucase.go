package usecase

import (
	"context"
	"time"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/domain/service"
	"clean_arch/usecase/in"
	"clean_arch/usecase/out"
)

// UserUsecase -
type UserUsecase interface {
	Fetch(ctx context.Context, num int64) ([]*out.User, error)
	GetByID(ctx context.Context, id int64) (*out.User, error)
	GetByName(ctx context.Context, name string) (*out.User, error)
	Create(ctx context.Context, u *in.PostUser) (int64, error)
	Update(ctx context.Context, u *in.PutUser) (*out.User, error)
	Delete(ctx context.Context, id int64) error
}

type userUsecase struct {
	repo       repository.UserRepository
	service    *service.UserService
	ctxTimeout time.Duration
}

// NewUserUseCase -
func NewUserUseCase(
	repo repository.UserRepository,
	service *service.UserService,
	timeout time.Duration,
) UserUsecase {
	return &userUsecase{
		repo:       repo,
		service:    service,
		ctxTimeout: timeout,
	}
}

func (u *userUsecase) Fetch(c context.Context, num int64) ([]*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	users, err := u.repo.Fetch(ctx, num)
	if err != nil {
		return nil, err
	}
	return toUsers(users), nil
}

func (u *userUsecase) GetByID(c context.Context, id int64) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &out.User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}, nil
}

func (u *userUsecase) GetByName(c context.Context, name string) (*out.User, error) {
	return nil, nil
}

func (u *userUsecase) Create(c context.Context, user *in.PostUser) (int64, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	return u.repo.Create(ctx, &model.User{
		Name: user.Name,
	})
}

func (u *userUsecase) Update(c context.Context, user *in.PutUser) (*out.User, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	usr, err := u.repo.Update(ctx, &model.User{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		return nil, err
	}
	return toUser(usr), nil
}

func (u *userUsecase) Delete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()
	return u.repo.Delete(ctx, id)
}

func (u *userUsecase) RegisterUser(c context.Context, name string) (int64, error) {
	if err := u.service.Duplicated(c, name); err != nil {
		return -1, err
	}
	user := model.NewUser(name)
	if id, err := u.repo.Create(c, user); err != nil {
		return id, err
	}
	return -1, nil
}

func toUser(user *model.User) *out.User {
	return &out.User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}
}
func toUsers(users []*model.User) []*out.User {
	res := make([]*out.User, len(users))
	for i, user := range users {
		res[i] = &out.User{
			ID:   user.GetID(),
			Name: user.GetName(),
		}
	}
	return res
}
