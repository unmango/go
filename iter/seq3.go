package iter

type Seq3[T, U, V any] func(yield func(T, U, V) bool)

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

func Singleton3[T, U, V any](t T, u U, v V) Seq3[T, U, V] {
	return func(yield func(T, U, V) bool) {
		_ = yield(t, u, v)
	}
}
