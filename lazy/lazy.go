package lazy

type Lazy[T any] func() T

func Of[T any](x T) Lazy[T] {
	return func() T {
		return x
	}
}

func Lift[L ~func() T, T any](l L) Lazy[T] {
	return Lazy[T](l)
}
