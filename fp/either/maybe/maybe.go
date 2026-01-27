package maybe

import (
	"errors"
)

type Maybe[T any] func() *T

var errNone = errors.New("None")

var ErrNone = errNone

func Some[T any](t T) Maybe[T] {
	return func() *T {
		return &t
	}
}

func None[T any]() Maybe[T] {
	return func() *T {
		return nil
	}
}

func Map[A, B any, M Maybe[A]](maybe M, fn func(A) B) Maybe[B] {
	return func() *B {
		if a := maybe(); a != nil {
			b := fn(*a)
			return &b
		}
		return nil
	}
}

func Bind[A, B any, M Maybe[A]](maybe M, fn func(A) Maybe[B]) Maybe[B] {
	return func() *B {
		if a := maybe(); a != nil {
			mb := fn(*a)
			return mb()
		}
		return nil
	}
}
