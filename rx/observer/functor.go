package observer

import "github.com/unmango/go/rx"

type functor[T, V any] struct {
	rx.Observer[T]
	Map func(V) T
}

// OnComplete implements rx.Observer.
func (m *functor[T, V]) OnComplete() {
	m.Observer.OnComplete()
}

// OnError implements rx.Observer.
func (m *functor[T, V]) OnError(err error) {
	m.Observer.OnError(err)
}

// OnNext implements rx.Observer.
func (m *functor[T, V]) OnNext(x V) {
	m.Observer.OnNext(m.Map(x))
}

func Map[T, V any](obs rx.Observer[T], fn func(V) T) rx.Observer[V] {
	return &functor[T, V]{obs, fn}
}
