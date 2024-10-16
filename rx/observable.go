package rx

type Subscription func()

func (s Subscription) Unsubscribe() {
	s()
}

type Subscriber[T any] interface {
	Next(T)
	Error(error)
	Complete()
}

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

type Observable[T any] interface {
	Subscribe(Observer[T]) Subscription
}

type Subject[T any] interface {
	Observer[T]
	Observable[T]
}
