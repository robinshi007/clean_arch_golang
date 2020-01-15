package respond

import (
	"bytes"
	"io"
	"net/http"

	Iserializer "clean_arch/domain/serializer"
	"clean_arch/domain/usecase/out"
	"clean_arch/endpoint/api"
)

// RespondMsgpack -
type RespondMsgpack struct {
	srz Iserializer.Serializer
}

// OK - write msgpack response format
func (r *RespondMsgpack) OK(w http.ResponseWriter, payload interface{}) {
	response, _ := r.Encode(api.NewResponse(payload))

	w.Header().Set("Content-Type", "application/x-msgpack")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Created - write msgpack response format
func (r *RespondMsgpack) Created(w http.ResponseWriter, payload interface{}) {
	response, _ := r.Encode(api.NewResponse(payload))

	w.Header().Set("Content-Type", "application/x-msgpack")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (r *RespondMsgpack) respondError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := r.srz.Encode(payload)
	w.Header().Set("Content-Type", "application/x-msgpack")
	w.WriteHeader(code)
	w.Write(response)
}

// Error - return error message
func (r *RespondMsgpack) Error(w http.ResponseWriter, err error) {
	code := GetErrorCode(err)
	r.respondError(w, out.GetHTTPStatus(code), api.NewErrorResponse(code))
}

// GraphQLError - return error message
func (r *RespondMsgpack) GraphQLError(w http.ResponseWriter, message string, path string) {
	r.respondError(w, http.StatusOK, api.NewGraphQLErrorResponse(message, path))
}

// Decode -
func (r *RespondMsgpack) Decode(input io.Reader, v interface{}) error {
	// convert io.Reader to []byte
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)
	return r.srz.Decode(buf.Bytes(), v)
}

// Encode -
func (r *RespondMsgpack) Encode(input interface{}) ([]byte, error) {
	return r.srz.Encode(input)
}
