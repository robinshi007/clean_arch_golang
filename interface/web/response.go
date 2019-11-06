package web

import (
	"clean_arch/usecase/output"
)

// Response - defined response json format
type Response struct {
	Success  bool            `json:"success"`
	Messages []string        `json:"messages"`
	Data     interface{}     `json:"data"`
	Errors   []*output.Error `json:"errors"`
}

// NewErrorResponse -
func NewErrorResponse(code string) *Response {
	return &Response{
		Success: false,
		Errors: []*output.Error{
			output.NewError(code),
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
		Errors:   []*output.Error{},
	}
}
