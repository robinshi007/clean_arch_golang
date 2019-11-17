package respond

import (
	"clean_arch/adapter/serializer"
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
	case "xml", "XML":
		return &RespondXML{
			srz: &serializer.XML{},
		}
	default:
		return &RespondJSON{
			srz: &serializer.JSON{},
		}
	}
}
