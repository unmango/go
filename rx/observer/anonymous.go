package observer

import "github.com/unmango/go/rx"

type Anonymous[T any] struct {
	Complete func()
	Error    func(error)
	Next     func(T)
}

// OnComplete implements rx.Observer.
func (a Anonymous[T]) OnComplete() {
	if a.Complete != nil {
		a.Complete()
	}
}

// OnError implements rx.Observer.
func (a Anonymous[T]) OnError(err error) {
	if a.Error != nil {
		a.Error(err)
	}
}

// OnNext implements rx.Observer.
func (a Anonymous[T]) OnNext(x T) {
	if a.Next != nil {
		a.Next(x)
	}
}

func Lift[F rx.Anonymous[T], T any](fn F) rx.Observer[T] {
	observer := &Anonymous[T]{}
	if c, ok := any(fn).(func()); ok {
		observer.Complete = c
	}
	if e, ok := any(fn).(func(error)); ok {
		observer.Error = e
	}
	if n, ok := any(fn).(func(T)); ok {
		observer.Next = n
	}

	return observer
}

func Subscribe[F rx.Anonymous[T], T any](
	observable rx.Observable[T], fn F,
) rx.Subscription {
	return observable.Subscribe(Lift[F, T](fn))
}
