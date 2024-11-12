package lazy

import "sync"

type Lazy[T any] func() T

type Lazy2[T, V any] func() (T, V)

type Lazy3[T, U, V any] func() (T, U, V)

func Of[T any](x T) Lazy[T] {
	return func() T {
		return x
	}
}

func Lift[L ~func() T, T any](l L) Lazy[T] {
	return Lazy[T](l)
}

func Once[L ~func() T, T any](l L) Lazy[T] {
	return sync.OnceValue(l)
}

func Once2[L ~func() (T, V), T, V any](l L) Lazy2[T, V] {
	return sync.OnceValues(l)
}
