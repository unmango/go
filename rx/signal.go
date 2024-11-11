package rx

type Signaler[T any] interface {
	Get() T
}

type Signal[T any] interface {
	Signaler[T]
	Set(T)
	Subscribe(func(T))
}
