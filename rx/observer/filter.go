package observer

import "github.com/unmango/go/rx"

type filter[T any] struct {
	src  rx.Observer[T]
	pred func(T) bool
}

// OnComplete implements rx.Observer.
func (f filter[T]) OnComplete() {
	f.src.OnComplete()
}

// OnError implements rx.Observer.
func (f filter[T]) OnError(err error) {
	f.src.OnError(err)
}

// OnNext implements rx.Observer.
func (f filter[T]) OnNext(next T) {
	if f.pred(next) {
		f.src.OnNext(next)
	}
}

func Filter[T any](obs rx.Observer[T], predicate func(T) bool) rx.Observer[T] {
	return filter[T]{obs, predicate}
}
