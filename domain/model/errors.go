package model

import "errors"

var (
	// ErrEntityNotFound - entity is not found in database
	ErrEntityNotFound = errors.New("requested item is not found")
	// ErrEntityUniqueConflict - entity cannot be created or updated due to unique conflict
	ErrEntityUniqueConflict = errors.New("requested item is already exist")
	// ErrEntityBadInput - entity cannot be created or updated due to bad input
	ErrEntityBadInput = errors.New("input params is not valid")
	// ErrEntityNotChanged - entity is not changed when updating
	ErrEntityNotChanged = errors.New("requested item is not changed")

	// ErrAccountNotExist -
	ErrAccountNotExist = errors.New("requested account is not exist")
	// ErrAccountNotMatch -
	ErrAccountNotMatch = errors.New("account email or password is not correct")
)
