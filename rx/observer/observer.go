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

type Anonymous2[K, V any] struct {
	Complete func()
	Error    func(error)
	Next     func(K, V)
}

// OnComplete implements rx.Observer.
func (a Anonymous2[K, V]) OnComplete() {
	if a.Complete != nil {
		a.Complete()
	}
}

// OnError implements rx.Observer.
func (a Anonymous2[K, V]) OnError(err error) {
	if a.Error != nil {
		a.Error(err)
	}
}

// OnNext implements rx.Observer.
func (a Anonymous2[K, V]) OnNext(k K, v V) {
	if a.Next != nil {
		a.Next(k, v)
	}
}

func Lift[F rx.AnonymousObserver[T], T any](fn F) rx.Observer[T] {
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

func Lift2[F rx.AnonymousObserver2[K, V], K, V any](fn F) rx.Observer2[K, V] {
	observer := &Anonymous2[K, V]{}
	if c, ok := any(fn).(func()); ok {
		observer.Complete = c
	}
	if e, ok := any(fn).(func(error)); ok {
		observer.Error = e
	}
	if n, ok := any(fn).(func(K, V)); ok {
		observer.Next = n
	}

	return observer
}

func Subscribe[O rx.AnonymousObserver[T], T any](
	observable rx.Observable[T], fn O,
) rx.Subscription {
	return observable.Subscribe(Lift[O, T](fn))
}

func WithNext[T any](obs rx.Observer[T], next func(T)) rx.Observer[T] {
	return Anonymous[T]{
		Complete: obs.OnComplete,
		Error:    obs.OnError,
		Next:     next,
	}
}
