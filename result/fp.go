package result

func Id[T any](r Result[T]) Result[T] {
	return r
}

func Map[T, V any](r Result[T], fn func(T) V) Result[V] {
	return func() (V, error) {
		if t, err := r(); err != nil {
			return zero[V](), err
		} else {
			return fn(t), nil
		}
	}
}

func Bind[T, V any](r Result[T], fn func(T) Result[V]) Result[V] {
	return func() (V, error) {
		if t, err := r(); err != nil {
			return zero[V](), err
		} else {
			return fn(t)()
		}
	}
}
