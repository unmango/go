package slices

import "github.com/unmango/go/iter"

func Collect2[T, U any](seq iter.Seq2[T, U]) (a []T, b []U) {
	for k, v := range seq {
		a = append(a, k)
		b = append(b, v)
	}

	return
}

func Collect3[T, U, V any](seq iter.Seq3[T, U, V]) (a []T, b []U, c []V) {
	seq(func(t T, u U, v V) bool {
		a = append(a, t)
		b = append(b, u)
		c = append(c, v)

		return true
	})

	return
}
