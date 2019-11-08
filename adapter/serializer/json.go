package serializer

import (
	"encoding/json"
	"fmt"
)

// JSON -
type JSON struct{}

// Decode -
func (r *JSON) Decode(input []byte, v interface{}) error {
	if err := json.Unmarshal(input, v); err != nil {
		return fmt.Errorf("serializer.JSON.Decode: %w", err)
	}
	return nil
}

// Encode -
func (r *JSON) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("serializer.JSON.Encode: %w", err)
	}
	return rawMsg, nil
}
