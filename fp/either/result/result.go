package result

import (
	"errors"
	"fmt"
)

type Result[T any] func() (T, error)

func Ok[T any](t T) Result[T] {
	return func() (T, error) {
		return t, nil
	}
}

func Error[T any](err error) Result[T] {
	return func() (T, error) {
		return *new(T), err
	}
}

func ErrorString[T any](text string) Result[T] {
	return Error[T](errors.New(text))
}

func Errorf[T any](format string, a ...any) Result[T] {
	return Error[T](fmt.Errorf(format, a...))
}

func Map[A, B any, R Result[A]](result R, fn func(A) B) Result[B] {
	return func() (B, error) {
		if a, err := result(); err != nil {
			return *new(B), err
		} else {
			return fn(a), nil
		}
	}
}

func Bind[A, B any, R Result[A]](result R, fn func(A) Result[B]) Result[B] {
	return func() (B, error) {
		if a, err := result(); err != nil {
			return *new(B), err
		} else {
			rb := fn(a)
			return rb()
		}
	}
}

type Result2[T, V any] func() (T, V, error)

func Ok2[T, V any](t T, v V) Result2[T, V] {
	return func() (T, V, error) {
		return t, v, nil
	}
}

func Error2[T, V any](err error) Result2[T, V] {
	return func() (T, V, error) {
		return *new(T), *new(V), err
	}
}

func Map2[TA, VA, TB, VB any, R Result2[TA, VA]](result R, fn func(TA, VA) (TB, VB)) Result2[TB, VB] {
	return func() (TB, VB, error) {
		if t, v, err := result(); err != nil {
			return *new(TB), *new(VB), err
		} else {
			tb, vb := fn(t, v)
			return tb, vb, nil
		}
	}
}

func Bind2[TA, VA, TB, VB any, R Result2[TA, VA]](result R, fn func(TA, VA) Result2[TB, VB]) Result2[TB, VB] {
	return func() (TB, VB, error) {
		if t, v, err := result(); err != nil {
			return *new(TB), *new(VB), err
		} else {
			return fn(t, v)()
		}
	}
}
