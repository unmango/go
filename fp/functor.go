package fp

type Functor[T, V any] interface {
	func(T) V
}
