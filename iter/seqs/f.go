package seqs

import (
	"github.com/unmango/go/iter"
	"github.com/unmango/go/result"
)

func Append[V any](seq iter.Seq[V], v V) iter.Seq[V] {
	ref := v
	return func(yield func(V) bool) {
		seq(yield)
		yield(ref)
	}
}

func FilterR[V any](seq iter.Seq[result.Result[V]]) iter.Seq[V] {
	filtered := Filter(seq, func(v result.Result[V]) bool {
		return v.IsOk()
	})

	return Map(filtered, func(v result.Result[V]) V {
		return v.Value()
	})
}

func Filter[V any](seq iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(v V) bool {
			if predicate(v) {
				return yield(v)
			}

			return true
		})
	}
}

func Map[V, X any](seq iter.Seq[V], f func(V) X) iter.Seq[X] {
	return func(yield func(X) bool) {
		seq(func(v V) bool {
			return yield(f(v))
		})
	}
}
