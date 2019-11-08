package respond

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"

	"clean_arch/domain/model"
	Iserializer "clean_arch/domain/serializer"
	"clean_arch/domain/usecase/out"
	"clean_arch/interface/api"
)

// RespondXML -
type RespondXML struct {
	srz Iserializer.Serializer
}

// OK - write xml response format
func (r *RespondXML) OK(w http.ResponseWriter, payload interface{}) {
	response, _ := r.Encode(api.NewResponse(payload))

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Created - write xml response format
func (r *RespondXML) Created(w http.ResponseWriter, payload interface{}) {
	response, _ := r.Encode(api.NewResponse(payload))

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (r *RespondXML) respondError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := r.srz.Encode(payload)
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(code)
	w.Write(response)
}

// Error - return error message
func (r *RespondXML) Error(w http.ResponseWriter, err error) {
	var code string
	if strings.HasPrefix(err.Error(), "pq:") {
		code = "1103"
	} else {
		switch {
		case errors.Is(err, model.ErrEntityBadInput):
			code = "1101"
		case errors.Is(err, model.ErrEntityNotFound):
			code = "1104"
		case errors.Is(err, model.ErrEntityUniqueConflict):
			code = "1104"
		default:
			code = "1103"
		}
	}
	r.respondError(w, out.GetHTTPStatus(code), api.NewErrorResponse(code))
}

// Decode -
func (r *RespondXML) Decode(input io.Reader, v interface{}) error {
	// convert io.Reader to []byte
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)
	return r.srz.Decode(buf.Bytes(), v)
}

// Encode -
func (r *RespondXML) Encode(input interface{}) ([]byte, error) {
	return r.srz.Encode(input)
}
