package api

import (
	"clean_arch/domain/usecase/out"
	"encoding/xml"
)

// Response - defined response json format
type Response struct {
	XMLName  xml.Name     `json:"-" xml:"response"`
	Success  bool         `json:"success" xml:"success"`
	Messages []string     `json:"messages" xml:"messages"`
	Data     interface{}  `json:"data" xml:"data>data"`
	Errors   []*out.Error `json:"errors" xml:"errors"`
}

// NewErrorResponse -
func NewErrorResponse(code string) *Response {
	return &Response{
		Success: false,
		Errors: []*out.Error{
			out.GetError(code),
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
