package obsv

import "github.com/unmango/go/rx"

func New[T any](func(rx.Subscriber[T])) rx.Observable[T] {
	return &subject[T]{}
}
