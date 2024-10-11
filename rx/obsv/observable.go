package obs

import (
	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/subject"
)

type Option[T any] func()

func New[T any](options ...Option[T]) rx.Observable[T] {
	return subject.New[T]()
}
