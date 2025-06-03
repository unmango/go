package json

import "encoding/json"

var DefaultCodec = Codec{}

// Codec implements https://pkg.go.dev/encoding/json
type Codec struct{}

func (Codec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (Codec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
