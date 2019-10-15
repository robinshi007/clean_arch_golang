package usecase

import (
	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/domain/service"
	"clean_arch/usecase/in"
	"clean_arch/usecase/out"
)

// UserUsecase -
type UserUsecase interface {
	Fetch(num int64) ([]*out.User, error)
	GetByID(id int64) (*out.User, error)
	GetByName(name string) (*out.User, error)
	Create(u *in.PostUser) (int64, error)
	Update(u *in.PutUser) (*out.User, error)
	Delete(id int64) (bool, error)
}

type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

// NewUserUseCase -
func NewUserUseCase(repo repository.UserRepository, service *service.UserService) UserUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userUsecase) Fetch(num int64) ([]*out.User, error) {
	users, err := u.repo.Fetch(num)
	if err != nil {
		return nil, err
	}
	return toUsers(users), nil
}

func (u *userUsecase) GetByID(id int64) (*out.User, error) {
	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &out.User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}, nil
}

func (u *userUsecase) GetByName(name string) (*out.User, error) {
	return nil, nil
}

func (u *userUsecase) Create(user *in.PostUser) (int64, error) {
	return u.repo.Create(&model.User{
		Name: user.Name,
	})
}

func (u *userUsecase) Update(user *in.PutUser) (*out.User, error) {
	usr, err := u.repo.Update(&model.User{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		return nil, err
	}
	return toUser(usr), nil
}

func (u *userUsecase) Delete(id int64) (bool, error) {
	return u.repo.Delete(id)
}

func (u *userUsecase) RegisterUser(name string) (int64, error) {
	if err := u.service.Duplicated(name); err != nil {
		return -1, err
	}
	user := model.NewUser(name)
	if id, err := u.repo.Create(user); err != nil {
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
