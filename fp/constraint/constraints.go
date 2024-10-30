package constraint

type ContraMap[A, B any] interface {
	~func(B) A
}

type Foldable[A, B any, M Monoid[A]] interface {
	~func(M, B) A
}

type Functor[A, B any, HKTA Type[A], HKTB Type[B], M Map[A, B]] interface {
	~func(HKTA, M) HKTB
}

type HKT[T, V any] interface {
	Type[T]
	~func(V)
}

type Identity[T any] interface {
	~func() T
}

type Map[A, B any] interface {
	~func(A) B
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
