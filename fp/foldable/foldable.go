package foldable

import (
	"github.com/unmango/go/fp/constraint"
)

type Foldable[
	T, V any,
	M constraint.Monoid[T],
] interface {
	Fold(M, V) T
}
