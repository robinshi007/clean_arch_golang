package model

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrNotFound            = errors.New("Your requested item is not found")
	ErrConflict            = errors.New("Your item already exist")
	ErrBadParamInput       = errors.New("Given Param is not valid")
)
