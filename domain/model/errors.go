package model

import "errors"

var (
	// ErrEntityBadInput - entity cannot be created or updated due to bad input
	ErrEntityBadInput = errors.New("input params is not valid")
	// ErrEntityNotFound - entity is not found in database
	ErrEntityNotFound = errors.New("requested item is not found")
	// ErrEntityUniqueConflict - entity cannot be created or updated due to unique conflict
	ErrEntityUniqueConflict = errors.New("requested item is already exist")
	// ErrEntityNotChanged - entity is not changed when updating
	ErrEntityNotChanged = errors.New("requested item is not changed")
	// ErrInternalServerError - internal server error
	ErrInternalServerError = errors.New("internal server error")

	// ErrAuthNotMatch -
	ErrAuthNotMatch = errors.New("login email or password is not correct")
	// ErrTokenIsExpired -
	ErrTokenExpired = errors.New("token is expired")
	// ErrTokenIsInvalid -
	ErrTokenIsInvalid = errors.New("token is invalid")
	// ErrTokenIsInvalid -
	ErrActionNotAllowed = errors.New("action is not allowed")
)
