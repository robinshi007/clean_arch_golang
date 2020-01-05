package model

import "errors"

var (
	// ErrRouteNotFound -
	ErrRouteNotFound = errors.New("route not found")
	// ErrMethodNotAllowed -
	ErrMethodNotAllowed = errors.New("http method not allowed")

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
	// ErrTokenEmpty -
	ErrTokenEmpty = errors.New("token is empty")
	// ErrTokenExpired -
	ErrTokenExpired = errors.New("token is expired")
	// ErrTokenInvalid -
	ErrTokenInvalid = errors.New("token is invalid")
	// ErrActionNotAllowed -
	ErrActionNotAllowed = errors.New("action is not allowed")
)
