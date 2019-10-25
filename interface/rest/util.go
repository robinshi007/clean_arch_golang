package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"clean_arch/domain/model"
	"clean_arch/usecase/output"
)

// RespondOK write json response format
func RespondOK(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(NewResponse(payload))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// RespondCreated write json response format
func RespondCreated(w http.ResponseWriter, payload interface{}) {
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

// RespondError return error message
func RespondError(w http.ResponseWriter, err error) {
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
	respondError(w, output.GetHTTPStatus(code), NewErrorResponse(code))
}
