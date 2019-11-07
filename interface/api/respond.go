package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"clean_arch/domain/model"
	"clean_arch/domain/usecase/out"
)

// Respond -
type Respond interface {
	OK(w http.ResponseWriter, payload interface{})
	Created(w http.ResponseWriter, payload interface{})
	Error(w http.ResponseWriter, err error)
}

// NewRespond -
func NewRespond(code string) Respond {
	switch code {
	case "json", "JSON":
		return &RespondJSON{}
	default:
		return &RespondJSON{}
	}
}

// RespondJSON -
type RespondJSON struct {
}

// OK - write json response format
func (r *RespondJSON) OK(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(NewResponse(payload))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Created - write json response format
func (r *RespondJSON) Created(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(NewResponse(payload))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func respondError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Error - return error message
func (r *RespondJSON) Error(w http.ResponseWriter, err error) {
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
	respondError(w, out.GetHTTPStatus(code), NewErrorResponse(code))
}
