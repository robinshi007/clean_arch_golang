package input

import (
	"gopkg.in/go-playground/validator.v9"

	"clean_arch/domain/model"
)

// PostUser -
type PostUser struct {
	Name string `validate:"required"`
}

// Validate -
func (u PostUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// PutUser -
type PutUser struct {
	User *model.User
	Name string `validate:"required"`
}

// Validate -
func (u PutUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
