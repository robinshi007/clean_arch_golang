package api

import (
	"io"
	"net/http"
)

// Responder -
// Endpoint level interface
type Responder interface {
	OK(w http.ResponseWriter, payload interface{})
	Created(w http.ResponseWriter, payload interface{})
	Error(w http.ResponseWriter, err error)
	GraphQLError(w http.ResponseWriter, message string, path string)
	Decode(input io.Reader, v interface{}) error
	Encode(input interface{}) ([]byte, error)
}
