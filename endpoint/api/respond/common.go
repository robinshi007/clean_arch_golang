package respond

import (
	"errors"
	"strings"

	"clean_arch/adapter/serializer"
	"clean_arch/domain/model"
	"clean_arch/endpoint/api"
)

// NewRespond -
func NewRespond(code string) api.Respond {
	switch code {
	case "json", "JSON":
		return &RespondJSON{
			srz: &serializer.JSON{},
		}
	case "msgpack", "Msgpack", "MsgPack":
		return &RespondMsgpack{
			srz: &serializer.Msgpack{},
		}
	default:
		return &RespondJSON{
			srz: &serializer.JSON{},
		}
	}
}

// GetErrorCode -
func GetErrorCode(err error) string {
	var code string
	if strings.HasPrefix(err.Error(), "pq: duplicate key value violates unique constraint") {
		code = "101"
	} else if strings.HasPrefix(err.Error(), "pq:") {
		code = "103"
	} else {
		switch {
		case errors.Is(err, model.ErrEntityBadInput):
			code = "101"
		case errors.Is(err, model.ErrEntityNotFound):
			code = "102"
		case errors.Is(err, model.ErrEntityNotChanged):
			code = "103"
		case errors.Is(err, model.ErrEntityUniqueConflict):
			code = "104"
		case errors.Is(err, model.ErrInternalServerError):
			code = "105"
		case errors.Is(err, model.ErrRouteNotFound):
			code = "106"
		case errors.Is(err, model.ErrMethodNotAllowed):
			code = "107"
		case errors.Is(err, model.ErrAuthNotMatch):
			code = "201"
		case errors.Is(err, model.ErrTokenExpired):
			code = "202"
		case errors.Is(err, model.ErrTokenEmpty):
			code = "203"
		case errors.Is(err, model.ErrTokenInvalid):
			code = "204"
		case errors.Is(err, model.ErrActionNotAllowed):
			code = "205"
		default:
			code = "103"
		}
	}
	return code
}
