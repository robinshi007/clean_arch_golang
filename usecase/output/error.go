package output

import (
	"net/http"
)

// Error -
type Error struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewBadReqeustError -
func NewBadReqeustError() *Error {
	return &Error{
		Status:  http.StatusBadRequest,
		Code:    "1101",
		Message: "Bad Request",
	}
}

// NewNotFoundError -
func NewNotFoundError() *Error {
	return &Error{
		Status:  http.StatusNotFound,
		Code:    "1102",
		Message: "Not Found",
	}
}

// NewInternalServerError -
func NewInternalServerError() *Error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Code:    "1103",
		Message: "Internal Server Error",
	}
}

// NewConflictError -
func NewConflictError() *Error {
	return &Error{
		Status:  http.StatusConflict,
		Code:    "1104",
		Message: "Conflict",
	}
}

// NewUnauthorizedError -
func NewUnauthorizedError() *Error {
	return &Error{
		Status:  http.StatusUnauthorized,
		Code:    "1105",
		Message: "Unauthorized",
	}
}

// NewForbiddenError -
func NewForbiddenError() *Error {
	return &Error{
		Status:  http.StatusForbidden,
		Code:    "1106",
		Message: "Forbidden",
	}
}

// NewRequestTimeoutError -
func NewRequestTimeoutError() *Error {
	return &Error{
		Status:  http.StatusRequestTimeout,
		Code:    "1107",
		Message: "Request Timeout",
	}
}

// NewUnsupportedMediaTypeError -
func NewUnsupportedMediaTypeError() *Error {
	return &Error{
		Status:  http.StatusUnsupportedMediaType,
		Code:    "1108",
		Message: "Unsupported Media Type",
	}
}

// NewError -
func NewError(code string) *Error {
	switch code {
	case "1101":
		return NewBadReqeustError()
	case "1102":
		return NewNotFoundError()
	case "1103":
		return NewInternalServerError()
	case "1104":
		return NewConflictError()
	case "1105":
		return NewUnauthorizedError()
	case "1106":
		return NewForbiddenError()
	case "1107":
		return NewRequestTimeoutError()
	case "1108":
		return NewUnsupportedMediaTypeError()
	default:
		return NewInternalServerError()
	}
}

// GetHTTPStatus -
func GetHTTPStatus(code string) int {
	switch code {
	case "1101":
		return http.StatusBadRequest
	case "1102":
		return http.StatusNotFound
	case "1103":
		return http.StatusInternalServerError
	case "1104":
		return http.StatusConflict
	case "1105":
		return http.StatusUnauthorized
	case "1106":
		return http.StatusForbidden
	case "1107":
		return http.StatusRequestTimeout
	case "1108":
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}
}
