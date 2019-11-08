package api

import (
	"io"
	"net/http"
)

// Respond -
// interface level interface
type Respond interface {
	OK(w http.ResponseWriter, payload interface{})
	Created(w http.ResponseWriter, payload interface{})
	Error(w http.ResponseWriter, err error)
	Decode(input io.Reader, v interface{}) error
	Encode(input interface{}) ([]byte, error)
}
