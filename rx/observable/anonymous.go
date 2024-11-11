package observable

import "github.com/unmango/go/rx"

type Anonymous[T any] func(rx.Observer[T]) rx.Subscription

// Subscribe implements rx.Observable.
func (a Anonymous[T]) Subscribe(obs rx.Observer[T]) rx.Subscription {
	return a(obs)
}
