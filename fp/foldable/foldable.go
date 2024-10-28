package foldable

import "github.com/unmango/go/fp/monoid"

type Foldable[T, V any] interface {
	Fold(monoid.Monoid[T], V) T
}
