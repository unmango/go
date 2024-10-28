package constraints

type Foldable[T, V any, M Monoid[T]] interface {
	~func(M, V) T
}

type Functor[T, V any] interface {
	~func(T) V
}

type Functor2[T, V, X, Y any] interface {
	~func(T, V) (X, Y)
}

type Functor3[T, U, V, X, Y, Z any] interface {
	~func(T, U, V) (X, Y, Z)
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
