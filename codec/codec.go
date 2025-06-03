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

type Decoder interface {
	Decode(v any) error
}

type Encoder interface {
	Encode(v any) error
}

type Marshaler interface {
	Marshal(v any) ([]byte, error)
}

type Unmarhsaler interface {
	Unmarshal(data []byte, v any) error
}
