package service

import (
	"fmt"

	"clean_arch/domain/repository"
)

// UserService -
type UserService struct {
	repo repository.UserRepository
}

// NewUserService -
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Duplicated -
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
