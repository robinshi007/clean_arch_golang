package out

import (
	"net/http"
)

// Error -
type Error struct {
	Status  int    `json:"-" xml:"-"`
	Code    string `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
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
		Message: "Internal Server Error",
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

// NewUnauthorizedError -
func NewUnauthorizedError() *Error {
	return &Error{
		Status:  http.StatusUnauthorized,
		Code:    "203",
		Message: "The Action is Unauthorized",
	}
}

// NewForbiddenError -
func NewForbiddenError() *Error {
	return &Error{
		Status:  http.StatusForbidden,
		Code:    "204",
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

// GetError -
func GetError(code string) *Error {
	switch code {
	// common error
	case "101":
		return NewBadReqeustError()
	case "102":
		return NewNotFoundError()
	case "103":
		return NewNotChangedError()
	case "104":
		return NewConflictError()
	case "105":
		return NewInternalServerError()

	// auth and permisson error
	case "201":
		return NewLoginError()
	case "202":
		return NewTokenExpiredError()
	case "203":
		return NewUnauthorizedError()
	case "204":
		return NewForbiddenError()

	// misc error
	case "901":
		return NewRequestTimeoutError()
	case "902":
		return NewUnsupportedMediaTypeError()

	default:
		return NewInternalServerError()
	}
}

// GetHTTPStatus -
func GetHTTPStatus(code string) int {
	switch code {
	case "101":
		return http.StatusBadRequest
	case "102":
		return http.StatusNotFound
	case "103":
		return http.StatusNotModified
	case "104":
		return http.StatusConflict
	case "105":
		return http.StatusInternalServerError

	case "201":
		return http.StatusUnauthorized
	case "202":
		return http.StatusUnauthorized
	case "203":
		return http.StatusUnauthorized
	case "204":
		return http.StatusForbidden

	case "901":
		return http.StatusRequestTimeout
	case "902":
		return http.StatusUnsupportedMediaType

	default:
		return http.StatusInternalServerError
	}
}
