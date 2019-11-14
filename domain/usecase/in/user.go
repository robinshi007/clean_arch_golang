package in

import (
	"gopkg.in/go-playground/validator.v9"

	"clean_arch/domain/model"
)

// NewUser -
type NewUser struct {
	Name string `validate:"required"`
}

// Validate -
func (u NewUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// EditUser -
type EditUser struct {
	User *model.User
	Name string `validate:"required"`
}

// Validate -
func (u EditUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// FetchUser -
type FetchUser struct {
	ID int64 `validate:"required"`
}

// Validate -
func (u FetchUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// FetchUserInput -
type FetchUserInput struct {
	ID string `json:"id"`
}

// EditUserInput -
type EditUserInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
