package slices

import (
	"fmt"

	"github.com/unmango/go/iter"
)

func All2[T, U any](t []T, u []U) iter.Seq3[int, T, U] {
	if len(t) != len(u) {
		panic(fmt.Sprintf("invalid source elements: t=%d u=%d", len(t), len(u)))
	}

	return func(yield func(int, T, U) bool) {
		for i := range t {
			if !yield(i, t[i], u[i]) {
				return
			}
		}
	}
}

func AppendSeq2[T, U any](seq iter.Seq2[T, U], a T, b U) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		cont := true
		seq(func(t T, u U) bool {
			cont = yield(t, u)
			return cont
		})

		if cont {
			_ = yield(a, b)
		}
	}
}

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

func CompactSeq[E comparable](seq iter.Seq[E]) []E {
	return Compact(Collect(seq))
}

func CompactSeqFunc[E any](seq iter.Seq[E], eq func(E, E) bool) []E {
	return CompactFunc(Collect(seq), eq)
}

func Zip[A, B any](as []A, bs []B) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		for i := 0; i < len(as) && i < len(bs); i++ {
			if !yield(as[i], bs[i]) {
				return
			}
		}
	}
}

func Values3[A, B, C any](as []A, bs []B, cs []C) iter.Seq3[A, B, C] {
	return func(yield func(A, B, C) bool) {
		for i := 0; i < len(as) && i < len(bs) && i < len(cs); i++ {
			if !yield(as[i], bs[i], cs[i]) {
				return
			}
		}
	}
}
