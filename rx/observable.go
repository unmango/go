package rx

type Subscription func()

func (s Subscription) Unsubscribe() { s() }

type NextObserver[T any] interface {
	OnNext(T)
}

type NextObserver2[K, V any] interface {
	OnNext(K, V)
}

type ErrorObserver interface {
	OnError(error)
}

type CompletionObserver interface {
	OnComplete()
}

type Observer[T any] interface {
	NextObserver[T]
	ErrorObserver
	CompletionObserver
}

type Observer2[K, V any] interface {
	NextObserver2[K, V]
	ErrorObserver
	CompletionObserver
}

type AnonymousObserver[T any] interface {
	~func(T) | ~func(error) | ~func()
}

type AnonymousObserver2[K, V any] interface {
	~func(K, V) | ~func(error) | ~func()
}

type Observable[T any] interface {
	Subscribe(Observer[T]) Subscription
}

type Observable2[K, V any] interface {
	Subscribe(Observer2[K, V]) Subscription
}

type AnonymousObservable[T any] interface {
	~func(Observer[T]) Subscription
}

type AnonymousObservable2[K, V any] interface {
	~func(Observer2[K, V]) Subscription
}

type Subject[T any] interface {
	Observer[T]
	Observable[T]
}

type Subject2[K, V any] interface {
	Observer2[K, V]
	Observable2[K, V]
}
