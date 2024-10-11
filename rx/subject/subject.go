package subject

import (
	"github.com/unmango/go/iter"
	"github.com/unmango/go/iter/seqs"
	"github.com/unmango/go/rx"
)

type subject[T any] struct {
	subs iter.Seq[rx.Observer[T]]
}

type Option[T any] func(*subject[T])

func (option Option[T]) apply(s *subject[T]) {
	option(s)
}

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
	s.subs = seqs.Append(s.subs, observer)

	return func() {
		s.subs = seqs.Remove(s.subs, observer)
	}
}

func New[T any](options ...Option[T]) rx.Subject[T] {
	subject := &subject[T]{}
	for _, opt := range options {
		opt.apply(subject)
	}

	return subject
}

func WithAnonymous[T any](func(rx.Subscriber[T])) Option[T] {
	return func(s *subject[T]) {} // TODO
}
