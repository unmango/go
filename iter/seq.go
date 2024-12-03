package iter

import (
	"errors"
	"iter"
)

type (
	// Seq is an iterator over sequences of individual values. When called as seq(yield),
	// seq calls yield(v) for each value v in the sequence, stopping early if yield
	// returns false. See the iter package documentation for more details.
	// Will be replaced with a type alias when "generic type aliases" is a stable feature
	Seq[V any] iter.Seq[V]
)

func D[V any](seq Seq[V]) iter.Seq[V] {
	return iter.Seq[V](seq)
}

func U[V any](seq iter.Seq[V]) Seq[V] {
	return Seq[V](seq)
}

func Append[V any](seq Seq[V], v V) Seq[V] {
	return func(yield func(V) bool) {
		cont := true
		seq(func(v V) bool {
			cont = yield(v)
			return cont
		})

		if cont {
			_ = yield(v)
		}
	}
}

func Bind[V, X any](seq Seq[V], fn func(V) Seq[X]) Seq[X] {
	return func(yield func(X) bool) {
		for v := range seq {
			for x := range fn(v) {
				if !yield(x) {
					return
				}
			}
		}
	}
}

func Empty[V any]() Seq[V] {
	return func(yield func(V) bool) {}
}

func Filter[V any](seq Seq[V], predicate func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !predicate(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Head[V any](seq Seq[V]) (v V, err error) {
	for v = range seq {
		return
	}

	return v, errors.New("empty sequence")
}

func Flat[T Seq[V], V any](seq Seq[T]) Seq[V] {
	return func(yield func(V) bool) {
		for s := range seq {
			for v := range s {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func FlatMap[T Seq[V], V, X any](seq Seq[T], fn func(V) X) Seq[X] {
	return func(yield func(X) bool) {
		for s := range seq {
			for v := range s {
				if !yield(fn(v)) {
					return
				}
			}
		}
	}
}

func Fold[A, V any](seq Seq[V], folder func(A, V) A, initial A) (acc A) {
	acc = initial
	for v := range seq {
		acc = folder(acc, v)
	}

	return
}

func Map[V, X any](seq Seq[V], fn func(V) X) Seq[X] {
	return func(yield func(X) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func()) {
	return iter.Pull(iter.Seq[V](seq))
}

func Remove[V comparable](seq Seq[V], r V) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if v == r {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Singleton[V any](v V) Seq[V] {
	return func(yield func(V) bool) {
		_ = yield(v)
	}
}

func Skip[V any](seq Seq[V], n int) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if n > 0 {
				n--
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Take[V any](seq Seq[V], n int) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if n <= 0 {
				return
			}
			if !yield(v) {
				return
			}

			n--
		}
	}
}

func Values[V any](values ...V) Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}
