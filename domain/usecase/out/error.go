package out

import (
	"net/http"
)

// Error -
type Error struct {
	Status  int    `json:"-" msgpack:"-"`
	Code    string `json:"code" msgpack:"code"`
	Message string `json:"message" msgpack:"message"`
}

// GraphQLError -
type GraphQLError struct {
	Message string   `json:"message" msgpack:"message"`
	Path    []string `json:"path" msgpack:"path"`
}

// NewBadReqeustError -
func NewBadReqeustError() *Error {
	return &Error{
		Status:  http.StatusBadRequest,
		Code:    "101",
		Message: "Bad Request",
	}
}

// NewNotFoundError -
func NewNotFoundError() *Error {
	return &Error{
		Status:  http.StatusNotFound,
		Code:    "102",
		Message: "Entity Not Found",
	}
}

// NewNotChangedError -
func NewNotChangedError() *Error {
	return &Error{
		Status:  http.StatusNotModified,
		Code:    "103",
		Message: "Entity Not Changed",
	}
}

// NewConflictError -
func NewConflictError() *Error {
	return &Error{
		Status:  http.StatusConflict,
		Code:    "104",
		Message: "Entity Conflict",
	}
}

// NewInternalServerError -
func NewInternalServerError() *Error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Code:    "105",
		Message: "OHH, Internal Server Error",
	}
}

// NewRouteNotFound -
func NewRouteNotFound() *Error {
	return &Error{
		Status:  http.StatusNotFound,
		Code:    "106",
		Message: "Route Path Not Found",
	}
}

// NewMethodNotAllowed -
func NewMethodNotAllowed() *Error {
	return &Error{
		Status:  http.StatusMethodNotAllowed,
		Code:    "107",
		Message: "HTTP Method Not Allowed",
	}
}

// NewLoginError -
func NewLoginError() *Error {
	return &Error{
		Status:  http.StatusUnauthorized,
		Code:    "201",
		Message: "Account email or password is not correct",
	}
}

// NewTokenExpiredError -
func NewTokenExpiredError() *Error {
	return &Error{
		Status:  http.StatusUnauthorized,
		Code:    "202",
		Message: "Token Is Expired",
	}
}

// NewTokenEmptyError -
func NewTokenEmptyError() *Error {
	return &Error{
		Status:  http.StatusUnauthorized,
		Code:    "203",
		Message: "Token Is Empty",
	}
}

// NewUnauthorizedError -
func NewUnauthorizedError() *Error {
	return &Error{
		Status:  http.StatusUnauthorized,
		Code:    "204",
		Message: "The Action is Unauthorized",
	}
}

// NewForbiddenError -
func NewForbiddenError() *Error {
	return &Error{
		Status:  http.StatusForbidden,
		Code:    "205",
		Message: "The Action Is Not Allowed",
	}
}

// NewRequestTimeoutError -
func NewRequestTimeoutError() *Error {
	return &Error{
		Status:  http.StatusRequestTimeout,
		Code:    "1108",
		Message: "Request Timeout",
	}
}

// NewUnsupportedMediaTypeError -
func NewUnsupportedMediaTypeError() *Error {
	return &Error{
		Status:  http.StatusUnsupportedMediaType,
		Code:    "1109",
		Message: "Unsupported Media Type",
	}
}

// ErrorResponseMap -
var ErrorResponseMap = map[string]*Error{
	"101": NewBadReqeustError(),
	"102": NewNotFoundError(),
	"103": NewNotChangedError(),
	"104": NewConflictError(),
	"105": NewInternalServerError(),
	"106": NewRouteNotFound(),
	"107": NewMethodNotAllowed(),

	// auth and permisson error
	"201": NewLoginError(),
	"202": NewTokenExpiredError(),
	"203": NewTokenEmptyError(),
	"204": NewUnauthorizedError(),
	"205": NewForbiddenError(),

	// misc error
	"901": NewRequestTimeoutError(),
	"902": NewUnsupportedMediaTypeError(),
}

// GetErrorResponse -
func GetErrorResponse(code string) *Error {
	val, ok := ErrorResponseMap[code]
	if ok == true {
		return val
	}
	// default error response
	return ErrorResponseMap["105"]
}

// GetHTTPStatus -
func GetHTTPStatus(code string) int {
	return GetErrorResponse(code).Status
}
