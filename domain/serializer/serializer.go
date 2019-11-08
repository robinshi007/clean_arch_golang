package serializer

// Serializer -
type Serializer interface {
	Decode(input []byte, v interface{}) error
	Encode(input interface{}) ([]byte, error)
}
