package iter

import (
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

func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func()) {
	return iter.Pull(iter.Seq[V](seq))
}

func Empty[V any]() Seq[V] {
	return func(yield func(V) bool) {}
}

func Singleton[V any](v V) Seq[V] {
	return func(yield func(V) bool) {
		_ = yield(v)
	}
}
