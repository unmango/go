package iter

import (
	"iter"
)

type Seq2[K, V any] = iter.Seq2[K, V]

func Append2[K, V any](seq Seq2[K, V], k K, v V) Seq2[K, V] {
	if seq == nil {
		return Singleton2(k, v)
	}

	return func(yield func(K, V) bool) {
		for a, b := range seq {
			if !yield(a, b) {
				return
			}
		}

		_ = yield(k, v)
	}
}

func Bind2[KA, VA, KB, VB any](seq Seq2[KA, VA], fn func(KA, VA) Seq2[KB, VB]) Seq2[KB, VB] {
	return func(yield func(KB, VB) bool) {
		seq(func(k KA, v VA) bool {
			for kb, vb := range fn(k, v) {
				if !yield(kb, vb) {
					return false
				}
			}
			return true
		})
	}
}

func DropFirst2[T, U any](seq Seq2[T, U]) Seq[U] {
	return func(yield func(U) bool) {
		seq(func(_ T, u U) bool {
			return yield(u)
		})
	}
}

func DropLast2[T, U any](seq Seq2[T, U]) Seq[T] {
	return func(yield func(T) bool) {
		seq(func(t T, _ U) bool {
			return yield(t)
		})
	}
}

func Empty2[K, V any]() Seq2[K, V] {
	return func(yield func(K, V) bool) {}
}

func Fold2[A, K, V any](seq Seq2[K, V], folder func(A, K, V) A, initial A) (acc A) {
	acc = initial
	for k, v := range seq {
		acc = folder(acc, k, v)
	}

	return
}

func Filter2[K, V any](seq Seq2[K, V], predicate func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !predicate(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func Head2[K, V any](seq Seq2[K, V]) (k K, v V) {
	for k, v = range Take2(seq, 1) {
		break
	}
	return
}

func Map2[K, V, X, Y any](seq Seq2[K, V], fn func(K, V) (X, Y)) Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		seq(func(k K, v V) bool {
			return yield(fn(k, v))
		})
	}
}

func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func()) {
	return iter.Pull2(seq)
}

func Singleton2[K, V any](k K, v V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		_ = yield(k, v)
	}
}

func Skip2[K, V any](seq Seq2[K, V], n int) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if n > 0 {
				n--
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func Take2[K, V any](seq Seq2[K, V], n int) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if n <= 0 {
				return
			}
			if !yield(k, v) {
				return
			}

			n--
		}
	}
}
