package in

import (
	"gopkg.in/go-playground/validator.v9"
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
	ID   int64  `validate:"required"`
	Name string `validate:"required"`
}

// Validate -
func (u PutUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
