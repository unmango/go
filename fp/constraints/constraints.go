package constraints

type Foldable[T, V any, M Monoid[T]] interface {
	~func(M, V) T
}

type Identity[T any] interface {
	~func() T
}

type Monoid[T any] interface {
	Semigroup[T]
	Identity[T]
}

type Semigroup[T any] interface {
	~func(T, T) T
}
