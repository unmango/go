package codec

import (
	"github.com/unmango/go/codec/googleproto"
	"github.com/unmango/go/codec/goyaml"
	"github.com/unmango/go/codec/json"
)

var (
	GoogleProto Codec = googleproto.DefaultCodec
	GoYaml      Codec = goyaml.DefaultCodec
	Json        Codec = json.DefaultCodec
)

type Codec interface {
	Marshaler
	Unmarhsaler
}

type Marshaler interface {
	Marshal(v any) ([]byte, error)
}

type Unmarhsaler interface {
	Unmarshal(data []byte, v any) error
}
