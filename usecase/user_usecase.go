package usecase

import (
	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/domain/service"
)

// UserUsecase -
type UserUsecase interface {
	ListUser() ([]*User, error)
	CreateUser(name string) error
}

type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

// NewUserUseCase -
func NewUserUseCase(repo repository.UserRepository, service *service.UserService) *UserUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userUsecase) ListUser() ([]*User, error) {
	users, err := u.repo.Fetch()
	if err != nil {
		return nil, err
	}
	return toUser(users), nil
}

func (u *userUsecase) RegisterUser(name string) error {
	if err := u.service.Duplicated(name); err != nil {
		return err
	}
	user := model.NewUser(name)
	if err := u.repo.Save(user); err != nil {
		return err
	}
	return nil
}

// User -
type User struct {
	ID    string
	Email string
}

func toUser(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID:    user.GetID(),
			Email: user.GetName(),
		}
	}
	return res
}
