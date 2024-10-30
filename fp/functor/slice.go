package functor

type sliceFunctor[A, B any] string

func (sliceFunctor[A, B]) Map(slice []A, f func(A) B) []B {
	acc := make([]B, len(slice))
	for i, a := range slice {
		acc[i] = f(a)
	}

	return acc
}
