package in

import (
	"strconv"

	"clean_arch/domain/model"

	"gopkg.in/go-playground/validator.v9"
)

// ToID -
func ToID(ID string) (int64, error) {
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return -1, model.ErrEntityBadInput
	}
	return id, nil
}

// Validate -
func Validate(v interface{}) error {
	validate := validator.New()
	return validate.Struct(v)
}
