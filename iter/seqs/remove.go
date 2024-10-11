package seqs

import "github.com/unmango/go/iter"

func Remove[T comparable](seq iter.Seq[T], v T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for x := range seq {
			if x != v {
				yield(x)
			}
		}
	}
}
