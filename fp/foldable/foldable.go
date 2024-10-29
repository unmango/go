package foldable

import (
	"github.com/unmango/go/fp/constraints"
)

type Foldable[
	T, V any,
	M constraints.Monoid[T],
] interface {
	Fold(M, V) T
}
