package api

import (
	"clean_arch/domain/usecase/out"
)

// Response - defined response json format
type Response struct {
	Success  bool         `json:"success"`
	Messages []string     `json:"messages"`
	Data     interface{}  `json:"data"`
	Errors   []*out.Error `json:"errors"`
}

// NewErrorResponse -
func NewErrorResponse(code string) *Response {
	return &Response{
		Success: false,
		Errors: []*out.Error{
			out.NewError(code),
		},
		Messages: []string{},
	}
}

// NewResponse -
func NewResponse(res interface{}) *Response {
	return &Response{
		Success:  true,
		Data:     res,
		Messages: []string{},
		Errors:   []*out.Error{},
	}
}
