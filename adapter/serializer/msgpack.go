package serializer

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

// Msgpack -
type Msgpack struct{}

// Decode -
func (r *Msgpack) Decode(input []byte, v interface{}) error {
	if err := msgpack.Unmarshal(input, v); err != nil {
		return fmt.Errorf("serializer.Msgpack.Decode: %w", err)
	}
	return nil
}

// Encode -
func (r *Msgpack) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("serializer.Msgpack.Encode: %w", err)
	}
	return rawMsg, nil
}
