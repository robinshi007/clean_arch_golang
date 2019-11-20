package in

import (
	"gopkg.in/go-playground/validator.v9"

	"clean_arch/domain/model"
)

// NewAccount -
type NewAccount struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// Validate -
func (u NewAccount) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// EditAccount -
type EditAccount struct {
	Account  *model.UserAccount
	Password string `validate:"required"`
}

// Validate -
func (u EditAccount) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// FetchAccount -
type FetchAccount struct {
	ID int64 `validate:"required"`
}

// Validate -
func (u FetchAccount) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// FetchAccountInput -
type FetchAccountInput struct {
	ID string `json:"id"`
}

// EditAccountInput -
type EditAccountInput struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
