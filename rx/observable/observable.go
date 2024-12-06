package observable

import (
	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/observer"
	"github.com/unmango/go/rx/subject"
)

type Anonymous[T any] func(rx.Observer[T]) rx.Subscription

// Subscribe implements rx.Observable.
func (a Anonymous[T]) Subscribe(obs rx.Observer[T]) rx.Subscription {
	return a(obs)
}

type Anonymous2[K, V any] func(rx.Observer2[K, V]) rx.Subscription

// Subscribe implements rx.Observable2.
func (a Anonymous2[K, V]) Subscribe(obs rx.Observer2[K, V]) rx.Subscription {
	return a(obs)
}

func Lift[O rx.AnonymousObservable[T], T any](obs O) rx.Observable[T] {
	return Anonymous[T](obs)
}

func Lift2[O rx.AnonymousObservable2[K, V], K, V any](obs O) rx.Observable2[K, V] {
	return Anonymous2[K, V](obs)
}

func New[T any](options ...subject.Option[T]) rx.Observable[T] {
	return subject.New(options...)
}

func Filter[T any](src rx.Observable[T], predicate func(T) bool) rx.Observable[T] {
	return Lift(func(obs rx.Observer[T]) rx.Subscription {
		return src.Subscribe(observer.Anonymous[T]{
			Complete: obs.OnComplete,
			Error:    obs.OnError,
			Next: func(t T) {
				if predicate(t) {
					obs.OnNext(t)
				}
			},
		})
	})
}

func Map[T, V any](src rx.Observable[T], project func(T) V) rx.Observable[V] {
	return Lift(func(obs rx.Observer[V]) rx.Subscription {
		return src.Subscribe(observer.Anonymous[T]{
			Complete: obs.OnComplete,
			Error:    obs.OnError,
			Next: func(t T) {
				obs.OnNext(project(t))
			},
		})
	})
}

// Bind is still TODO
func Bind[T, V any](src rx.Observable[T], fn func(T) rx.Observable[V]) rx.Observable[V] {
	return Lift(func(dest rx.Observer[V]) rx.Subscription {
		return src.Subscribe(observer.Anonymous[T]{
			Complete: dest.OnComplete,
			Error:    dest.OnError,
			Next: func(t T) {
				_ = fn(t).Subscribe(dest)
			},
		})
	})
}
