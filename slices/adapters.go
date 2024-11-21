package slices

import (
	"cmp"
	"slices"

	"github.com/unmango/go/iter"
)

func All[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return iter.U2(slices.All(s))
}

func AppendSeq[S ~[]E, E any](s S, seq iter.Seq[E]) S {
	return slices.AppendSeq(s, iter.D(seq))
}

func Backward[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return iter.U2(slices.Backward(s))
}

func Chunk[S ~[]E, E any](s S, n int) iter.Seq[S] {
	return iter.U(slices.Chunk(s, n))
}

func Collect[E any](s iter.Seq[E]) []E {
	return slices.Collect(iter.D(s))
}

func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E {
	return slices.Sorted(iter.D(seq))
}

func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return slices.SortedFunc(iter.D(seq), cmp)
}

func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E {
	return slices.SortedStableFunc(iter.D(seq), cmp)
}

func Values[S ~[]E, E any](s S) iter.Seq[E] {
	return iter.U(slices.Values(s))
}
