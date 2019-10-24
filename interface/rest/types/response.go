package types

import (
	"clean_arch/usecase/output"
	"encoding/json"
	"net/http"
	"strings"
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

// RespondWithJSON write json response format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(NewResponse(payload))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func respondWithError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError return error message
func RespondWithError(w http.ResponseWriter, code string, msg string) {
	if strings.HasPrefix(msg, "pq:") {
		respondWithError(w, output.GetHTTPStatus(code), map[string]string{"message": "Server Internal Error"})
	} else {
		respondWithError(w, output.GetHTTPStatus(code), NewErrorResponse(code))
	}
}
