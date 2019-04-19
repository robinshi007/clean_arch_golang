package service

import (
	"fmt"

	"github.com/robinshi007/goweb/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s *UserService) Duplicated(name string) error {
	user, err := s.repo.GetByName(name)
	if user != nil {
		return fmt.Errorf("%s already exists", name)
	}
	if err != nil {
		return err
	}
	return nil
}
