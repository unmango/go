package iter

type Seq3[T, U, V any] func(yield func(T, U, V) bool)

func Bind3[TA, UA, VA, TB, UB, VB any](seq Seq3[TA, UA, VA], fn func(TA, UA, VA) Seq3[TB, UB, VB]) Seq3[TB, UB, VB] {
	return func(yield func(TB, UB, VB) bool) {
		seq(func(t TA, u UA, v VA) bool {
			cont := true
			next := fn(t, u, v)
			next(func(tb TB, ub UB, vb VB) bool {
				cont = yield(tb, ub, vb)
				return cont
			})
			return cont
		})
	}
}

func DropLast3[T, U, V any](seq Seq3[T, U, V]) Seq2[T, U] {
	return func(yield func(T, U) bool) {
		seq(func(t T, u U, _ V) bool {
			return yield(t, u)
		})
	}
}

func DropMid3[T, U, V any](seq Seq3[T, U, V]) Seq2[T, V] {
	return func(yield func(T, V) bool) {
		seq(func(t T, _ U, v V) bool {
			return yield(t, v)
		})
	}
}

func DropFirst3[T, U, V any](seq Seq3[T, U, V]) Seq2[U, V] {
	return func(yield func(U, V) bool) {
		seq(func(_ T, u U, v V) bool {
			return yield(u, v)
		})
	}
}

func Empty3[T, U, V any]() Seq3[T, U, V] {
	return func(yield func(T, U, V) bool) {}
}

func Filter3[T, U, V any](seq Seq3[T, U, V], predicate func(T, U, V) bool) Seq3[T, U, V] {
	return func(yield func(T, U, V) bool) {
		seq(func(t T, u U, v V) bool {
			if predicate(t, u, v) {
				return yield(t, u, v)
			} else {
				return true
			}
		})
	}
}

func Fold3[A, T, U, V any](seq Seq3[T, U, V], folder func(A, T, U, V) A, initial A) (acc A) {
	acc = initial
	seq(func(t T, u U, v V) bool {
		acc = folder(acc, t, u, v)
		return true
	})

	return
}

func KeepFirst3[T, U, V any](seq Seq3[T, U, V]) Seq[T] {
	return func(yield func(T) bool) {
		seq(func(t T, _ U, _ V) bool {
			return yield(t)
		})
	}
}

func KeepMid3[T, U, V any](seq Seq3[T, U, V]) Seq[U] {
	return func(yield func(U) bool) {
		seq(func(_ T, u U, _ V) bool {
			return yield(u)
		})
	}
}

func KeepLast3[T, U, V any](seq Seq3[T, U, V]) Seq[V] {
	return func(yield func(V) bool) {
		seq(func(_ T, _ U, v V) bool {
			return yield(v)
		})
	}
}

func Map3[T, U, V, X, Y, Z any](seq Seq3[T, U, V], fn func(T, U, V) (X, Y, Z)) Seq3[X, Y, Z] {
	return func(yield func(X, Y, Z) bool) {
		seq(func(t T, u U, v V) bool {
			return yield(fn(t, u, v))
		})
	}
}

func Singleton3[T, U, V any](t T, u U, v V) Seq3[T, U, V] {
	return func(yield func(T, U, V) bool) {
		_ = yield(t, u, v)
	}
}
