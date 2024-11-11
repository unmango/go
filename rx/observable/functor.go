package observable

import (
	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/observer"
)

type functor[T, V any] struct {
	rx.Observable[T]
	Map func(T) V
}

// Subscribe implements rx.Observable.
func (m *functor[T, V]) Subscribe(obs rx.Observer[V]) rx.Subscription {
	return m.Observable.Subscribe(observer.Map(obs, m.Map))
}

func Map[T, V any](obs rx.Observable[T], fn func(T) V) rx.Observable[V] {
	return &functor[T, V]{obs, fn}
}
