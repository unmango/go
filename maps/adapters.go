package maps

import (
	"maps"

	"github.com/unmango/go/iter"
)

func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return iter.U2(maps.All(m))
}

func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	return maps.Collect(iter.D2(seq))
}

func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V]) {
	maps.Insert(m, iter.D2(seq))
}

func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return iter.U(maps.Keys(m))
}

func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return iter.U(maps.Values(m))
}
