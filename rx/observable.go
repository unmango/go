package rx

type Subscription func()

func (s Subscription) Unsubscribe() { s() }

type NextObserver[T any] interface {
	OnNext(T)
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

type AnonymousObserver[T any] interface {
	~func(T) | ~func(error) | ~func()
}

type Observable[T any] interface {
	Subscribe(Observer[T]) Subscription
}

type Subject[T any] interface {
	Observer[T]
	Observable[T]
}
