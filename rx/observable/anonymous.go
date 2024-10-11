package obs

import "github.com/unmango/go/rx"

type Anonymous[T any] func(rx.Subscriber[T])

// Subscribe implements rx.Observable.
func (a Anonymous[T]) Subscribe(observer rx.Observer[T]) rx.Subscription {
	a(&subscriber[T]{observer})
	return func() {}
}
