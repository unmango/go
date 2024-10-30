package semigroup

import "github.com/unmango/go/fp/constraint"

type Semigroup[T any] interface {
	Combine(T, T) T
}

type semigroup[T any] func(T, T) T

func (s semigroup[T]) Combine(a T, b T) T {
	return s(a, b)
}

func Lift[T any, S constraint.Semigroup[T]](fn S) Semigroup[T] {
	return semigroup[T](fn)
}
