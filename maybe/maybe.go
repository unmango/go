package maybe

import "errors"

type Maybe[T any] func() (T, error)

var ErrNone = errors.New("None")

func Ok[T any](v T) Maybe[T] {
	return func() (T, error) {
		return v, nil
	}
}

func None[T any]() (x T, err error) {
	err = ErrNone
	return
}

func IsOk[T any](m Maybe[T]) bool {
	return !IsNone(m)
}

func IsNone[T any](m Maybe[T]) bool {
	// TODO: Can this be done lazy?
	_, err := m()
	return err == ErrNone
}
