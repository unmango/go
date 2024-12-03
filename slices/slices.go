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

func Values2[T, U any](t []T, u []U) iter.Seq2[T, U] {
	if len(t) != len(u) {
		panic(fmt.Sprintf("invalid source elements: t=%d u=%d", len(t), len(u)))
	}

	return func(yield func(T, U) bool) {
		for i := range t {
			if !yield(t[i], u[i]) {
				return
			}
		}
	}
}

func Values3[T, U, V any](t []T, u []U, v []V) iter.Seq3[T, U, V] {
	if len(t) != len(u) || len(t) != len(v) || len(u) != len(v) {
		panic(fmt.Sprintf("invalid source elements: t=%d u=%d v=%d", len(t), len(u), len(v)))
	}

	return func(yield func(T, U, V) bool) {
		for i := range t {
			if !yield(t[i], u[i], v[i]) {
				return
			}
		}
	}
}
