package obs

import (
	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/subject"
)

func New[T any](options ...subject.Option[T]) rx.Observable[T] {
	return subject.New(options...)
}

func FromAnonymous[T any](observable func(rx.Subscriber[T])) rx.Observable[T] {
	return Anonymous[T](observable)
}
