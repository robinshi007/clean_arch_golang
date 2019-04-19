package util

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/robinshi007/goweb/domain/model"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// respondwithJSON write json response format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if strings.HasPrefix(msg, "pq:") {
		RespondWithJSON(w, code, map[string]string{"message": "Server Internal Error"})
	} else {
		RespondWithJSON(w, code, map[string]string{"message": msg})
	}
}
