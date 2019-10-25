package model

import "errors"

var (
	// ErrEntityNotFound - entity is not found in database
	ErrEntityNotFound = errors.New("requested item is not found")
	// ErrEntityUniqueConflict - entity cannot be created or updated due to unique conflict
	ErrEntityUniqueConflict = errors.New("requested item is already exist")
	// ErrEntityBadInput - entity cannot be created or updated due to bad input
	ErrEntityBadInput = errors.New("input params is not valid")
)
