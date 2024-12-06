package observer

import "github.com/unmango/go/rx"

type bind[T, V any] struct {
	src  rx.Observer[T]
	Bind func(rx.Observer[T]) V
}

// OnComplete implements rx.Observer.
func (b bind[T, V]) OnComplete() {
	b.src.OnComplete()
}

// OnError implements rx.Observer.
func (b bind[T, V]) OnError(err error) {
	b.src.OnError(err)
}

// OnNext implements rx.Observer.
func (b bind[T, V]) OnNext(next V) {
	panic("unimplemented")
}

func Bind[T, V any](obs rx.Observer[T], fn func(rx.Observer[T]) V) rx.Observer[V] {
	return bind[T, V]{obs, fn}
}
