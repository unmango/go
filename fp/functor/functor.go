package functor

import "github.com/unmango/go/fp/constraints"

type Functor[T, V any] interface {
	Map()
}

func Lift[T, V any, F constraints.Functor[T, V]]() {}
