package subject

import (
	"github.com/unmango/go/iter"
	"github.com/unmango/go/option"
	"github.com/unmango/go/rx"
)

type subject[T any] struct {
	subs iter.Seq[rx.Observer[T]]
}

type Option[T any] func(*subject[T])

// OnComplete implements rx.Subject.
func (s *subject[T]) OnComplete() {
	for sub := range s.subs {
		sub.OnComplete()
	}
}

// OnError implements rx.Subject.
func (s *subject[T]) OnError(err error) {
	for sub := range s.subs {
		sub.OnError(err)
	}
}

// OnNext implements rx.Subject.
func (s *subject[T]) OnNext(value T) {
	for sub := range s.subs {
		sub.OnNext(value)
	}
}

// Subscribe implements rx.Subject.
func (s *subject[T]) Subscribe(observer rx.Observer[T]) rx.Subscription {
	s.subs = iter.Append(s.subs, observer)

	return func() {
		s.subs = iter.Remove(s.subs, observer)
	}
}

func New[T any](options ...Option[T]) rx.Subject[T] {
	subs := iter.Empty[rx.Observer[T]]()
	subject := &subject[T]{subs}
	option.ApplyAll(subject, options)

	return subject
}
