package monoid

import (
	"github.com/unmango/go/fp/constraints"
	S "github.com/unmango/go/fp/semigroup"
)

type Monoid[T any] interface {
	S.Semigroup[T]
	Empty() T
}

type monoid[T any] struct {
	S.Semigroup[T]
	empty T
}

func (m monoid[T]) Empty() T {
	return m.empty
}

func From[T any](empty T, combine S.Semigroup[T]) Monoid[T] {
	return monoid[T]{combine, empty}
}

func Lift[T any, M constraints.Monoid[T]](monoid M) Monoid[T] {
	return monoid
}
