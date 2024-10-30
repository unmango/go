package constraint

type Foldable[A, B any, M Monoid[A]] interface {
	~func(M, B) A
}

type Functor[A, B any, HKTA Type[A], HKTB Type[B]] interface {
	~func(HKTA, func(A) B) HKTB
}

type HKT[T, V any] interface {
	Type[T]
	~func(V)
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

type Type[T any] interface {
	~func(T)
}
