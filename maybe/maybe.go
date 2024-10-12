package maybe

type Maybe[T any] func() (T, error)

type none struct{}

// Error implements error.
func (n none) Error() string { return "None" }

var ErrNone error = none{}

func Ok[T any](v T) Maybe[T] {
	return func() (T, error) {
		return v, nil
	}
}

func None() (any, error) {
	return nil, ErrNone
}

func IsOk[T any](m Maybe[T]) bool {
	return !IsNone(m)
}

func IsNone[T any](m Maybe[T]) bool {
	_, n := m()
	return n == ErrNone
}

func zero[T any]() T {
	var t T
	return t
}
