package goyaml

import "gopkg.in/yaml.v3"

var DefaultCodec = Codec{}

// Coded implements https://github.com/go-yaml/yaml
type Codec struct{}

func (Codec) Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Codec) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}
