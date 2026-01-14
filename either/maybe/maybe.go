package maybe

import (
	"errors"

	"github.com/unmango/go/either"
)

type Maybe[T any] = either.Either[T, error]

var errNone = errors.New("None")

var ErrNone = errNone

func Some[T any](v T) Maybe[T] {
	return either.Left[T, error](v)
}

func None[T any]() Maybe[T] {
	return either.Right[T](errNone)
}

func Map[A, B any, M Maybe[A]](maybe M, fn func(A) B) Maybe[B] {
	return func() (B, error) {
		if a, err := maybe(); err != nil {
			return *new(B), err
		} else {
			return fn(a), nil
		}
	}
}

func Bind[A, B any, M Maybe[A]](maybe M, fn func(A) Maybe[B]) Maybe[B] {
	return func() (B, error) {
		if a, err := maybe(); err != nil {
			return *new(B), err
		} else {
			return fn(a)()
		}
	}
}
