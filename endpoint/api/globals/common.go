package globals

import (
	"clean_arch/endpoint/api"
	"clean_arch/endpoint/api/respond"
	"clean_arch/registry"
)

var (
	// Respond -
	Respond api.Responder
)

// InitResponder -
func InitResponder() {
	Respond = respond.NewRespond(registry.Cfg.Serializer.Code)
}
