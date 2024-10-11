package obs

import (
	"github.com/unmango/go/iter"
	"github.com/unmango/go/iter/seqs"
	"github.com/unmango/go/rx"
)

type subject[T any] struct {
	subs iter.Seq[rx.Observer[T]]
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

func NewSubject[T any]() rx.Subject[T] {
	return &subject[T]{}
}
