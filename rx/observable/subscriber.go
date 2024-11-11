package observable

import "github.com/unmango/go/rx"

type subscriber[T any] struct{ rx.Observer[T] }

// Complete implements rx.Subscriber.
func (s *subscriber[T]) Complete() {
	s.OnComplete()
}

// Error implements rx.Subscriber.
func (s *subscriber[T]) Error(err error) {
	s.OnError(err)
}

// Next implements rx.Subscriber.
func (s *subscriber[T]) Next(value T) {
	s.OnNext(value)
}
