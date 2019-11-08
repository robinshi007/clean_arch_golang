package serializer

import (
	"encoding/xml"
	"fmt"
)

// XML -
type XML struct{}

// Decode -
func (r *XML) Decode(input []byte, v interface{}) error {
	if err := xml.Unmarshal(input, v); err != nil {
		return fmt.Errorf("serializer.XML.Decode: %w", err)
	}
	return nil
}

// Encode -
func (r *XML) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := xml.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("serializer.XML.Encode: %w", err)
	}
	return rawMsg, nil
}
