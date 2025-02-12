package maps

import (
	"maps"

	"github.com/unmango/go/iter"
)

func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return maps.All(m)
}

func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	return maps.Collect(seq)
}

func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V]) {
	maps.Insert(m, seq)
}

func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return maps.Keys(m)
}

func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return maps.Values(m)
}
