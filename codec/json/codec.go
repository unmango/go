package json

import (
	"encoding/json"
	"io"
)

var DefaultCodec = Codec{}

// Codec implements https://pkg.go.dev/encoding/json
type Codec struct{}

func (Codec) NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}

func (Codec) NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func (Codec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (Codec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
