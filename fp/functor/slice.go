package functor

import "github.com/unmango/go/fp/constraints"

type sliceFunctor[T, V any, F constraints.Functor[T, V]] string

func (sliceFunctor[T, V, F]) Map(x T, f F) V {
	return f(x)
}
