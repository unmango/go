package googleproto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

var DefaultCodec = Codec{}

type (
	MarshalOptions   = proto.MarshalOptions
	Message          = proto.Message
	UnmarshalOptions = proto.UnmarshalOptions
)

// Codec implements https://pkg.go.dev/google.golang.org/protobuf/proto
type Codec struct{}

func (c Codec) Marshal(v any) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return proto.Marshal(msg)
	} else {
		return nil, fmt.Errorf("not a proto.Message: %#v", v)
	}
}

func (c Codec) Unmarshal(data []byte, v any) error {
	if msg, ok := v.(proto.Message); ok {
		return proto.Unmarshal(data, msg)
	} else {
		return fmt.Errorf("not a proto.Message: %#v", v)
	}
}
