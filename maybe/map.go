package maybe

import (
	"errors"

	"github.com/unmango/go/monad"
)

func Map[T, V any, F monad.Functor[T, V]](m Maybe[T], f F) Maybe[V] {
	if v, err := m(); errors.Is(err, ErrNone) {
		return None
	} else {
		return Ok(f(v))
	}
}
