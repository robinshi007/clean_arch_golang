package respond

import (
	"bytes"
	"io"
	"net/http"

	Iserializer "clean_arch/domain/serializer"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api"
)

// RespondJSON -
type RespondJSON struct {
	srz Iserializer.Serializer
}

// OK - write json response format
func (r *RespondJSON) OK(w http.ResponseWriter, payload interface{}) {
	response, _ := r.Encode(api.NewResponse(payload))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Created - write json response format
func (r *RespondJSON) Created(w http.ResponseWriter, payload interface{}) {
	response, _ := r.Encode(api.NewResponse(payload))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (r *RespondJSON) respondError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := r.srz.Encode(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Error - return error message
func (r *RespondJSON) Error(w http.ResponseWriter, err error) {
	code := GetErrorCode(err)
	r.respondError(w, out.GetHTTPStatus(code), api.NewErrorResponse(code))
}

// Decode -
func (r *RespondJSON) Decode(input io.Reader, v interface{}) error {
	// convert io.Reader to []byte
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)
	return r.srz.Decode(buf.Bytes(), v)
}

// Encode -
func (r *RespondJSON) Encode(input interface{}) ([]byte, error) {
	return r.srz.Encode(input)
}
