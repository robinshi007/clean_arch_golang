package api

import (
	"clean_arch/domain/usecase/out"
)

// Response - defined response json format
type Response struct {
	Success  bool         `json:"success" msgpack:"success"`
	Messages []string     `json:"messages" msgpack:"messages"`
	Data     interface{}  `json:"data" msgpack:"data"`
	Errors   []*out.Error `json:"errors" msgpack:"errors"`
}

// GraphQLResponse -
type GraphQLResponse struct {
	Data   interface{}         `json:"data" msgpack:"data"`
	Errors []*out.GraphQLError `json:"errors" msgpack:"errors"`
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

// NewErrorResponse -
func NewErrorResponse(code string) *Response {
	return &Response{
		Success: false,
		Errors: []*out.Error{
			out.GetErrorResponse(code),
		},
		Messages: []string{},
	}
}

// NewGraphQLErrorResponse -
func NewGraphQLErrorResponse(message, path string) *GraphQLResponse {
	return &GraphQLResponse{
		Errors: []*out.GraphQLError{
			&out.GraphQLError{
				Message: message,
				Path:    []string{path},
			},
		},
		Data: nil,
	}
}
