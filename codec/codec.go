package codec

import "io"

type Any = Codec[any]

type Codec[T any] interface {
	Marshaler[T]

	NewDecoder(io.Reader) Decoder[T]
	NewEncoder(io.Writer) Encoder[T]
}

type Decoder[T any] interface {
	Decode(v T) error
}

type Encoder[T any] interface {
	Encode(v T) error
}

type Marshaler[T any] interface {
	Marshal(v T) ([]byte, error)
	Unmarshal(data []byte, v T) error
}

type decoder[T any] struct{ d Decoder[any] }

func (d decoder[T]) Decode(v T) error {
	return d.d.Decode(v)
}

type encoder[T any] struct{ e Encoder[any] }

func (e encoder[T]) Encode(v T) error {
	return e.e.Encode(v)
}

type cast[T any] struct{ c Codec[any] }

func Cast[T any, C Codec[any]](c C) Codec[T] {
	return cast[T]{c}
}

// Marshal implements [Codec].
func (c cast[T]) Marshal(v T) ([]byte, error) {
	return c.c.Marshal(v)
}

// NewDecoder implements [Codec].
func (c cast[T]) NewDecoder(r io.Reader) Decoder[T] {
	return decoder[T]{c.c.NewDecoder(r)}
}

// NewEncoder implements [Codec].
func (c cast[T]) NewEncoder(w io.Writer) Encoder[T] {
	return encoder[T]{c.c.NewEncoder(w)}
}

// Unmarshal implements [Codec].
func (c cast[T]) Unmarshal(b []byte, v T) error {
	return c.c.Unmarshal(b, v)
}
