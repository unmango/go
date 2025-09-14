package slices

import (
	"cmp"
	"iter"
	"slices"
)

func All[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return slices.All(s)
}

func AppendSeq[S ~[]E, E any](s S, seq iter.Seq[E]) S {
	return slices.AppendSeq(s, seq)
}

func Backward[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return slices.Backward(s)
}

func Compact[S ~[]E, E comparable](s S) []E {
	return slices.Compact(s)
}

func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) []E {
	return slices.CompactFunc(s, eq)
}

func Chunk[S ~[]E, E any](s S, n int) iter.Seq[S] {
	return slices.Chunk(s, n)
}

func Collect[E any](s iter.Seq[E]) []E {
	return slices.Collect(s)
}

func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E {
	return slices.Sorted(seq)
}

func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return slices.SortedFunc(seq, cmp)
}

func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return slices.SortedStableFunc(seq, cmp)
}

func Values[S ~[]E, E any](s S) iter.Seq[E] {
	return slices.Values(s)
}
