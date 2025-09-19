package codec

import (
	"github.com/unmango/go/codec/encodingjson"
	"github.com/unmango/go/codec/goccy"
	"github.com/unmango/go/codec/googleproto"
	"github.com/unmango/go/codec/goyaml"
)

var (
	Goccy        Codec = goccy.DefaultCodec
	GoogleProto  Codec = googleproto.DefaultCodec
	GoYaml       Codec = goyaml.DefaultCodec
	EncodingJson Codec = encodingjson.DefaultCodec

	// Defaults

	Json     = EncodingJson
	Yaml     = Goccy
	Protobuf = GoogleProto
)

type Codec interface {
	Marshaler
	Unmarshaler
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

type Unmarshaler interface {
	Unmarshal(data []byte, v any) error
}
