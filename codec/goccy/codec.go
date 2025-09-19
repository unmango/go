package goccy

import (
	"io"

	"github.com/goccy/go-yaml"
)

var DefaultCodec = Codec{}

type (
	Encoder = yaml.Encoder
	Decoder = yaml.Decoder
)

// Codec implements https://github.com/goccy/go-yaml
type Codec struct{}

func (Codec) NewDecoder(r io.Reader) *yaml.Decoder {
	return yaml.NewDecoder(r)
}

func (Codec) NewEncoder(w io.Writer) *yaml.Encoder {
	return yaml.NewEncoder(w)
}

func (Codec) Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Codec) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}
