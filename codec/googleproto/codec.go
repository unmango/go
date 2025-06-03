package googleproto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

var DefaultCodec = Codec{
	ErrorHandler: PanicHandler,
}

// Codec implements https://pkg.go.dev/google.golang.org/protobuf/proto
type Codec struct {
	ErrorHandler func(any) error
}

func (c Codec) Marshal(v any) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return proto.Marshal(msg)
	} else {
		return nil, c.ErrorHandler(v)
	}
}

func (c Codec) Unmarshal(data []byte, v any) error {
	if msg, ok := v.(proto.Message); ok {
		return proto.Unmarshal(data, msg)
	} else {
		return c.ErrorHandler(v)
	}
}

func PanicHandler(v any) error {
	panic(fmt.Sprintf("not a proto.Message: %v", v))
}
