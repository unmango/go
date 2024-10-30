package constraints

type Foldable[A, B any, M Monoid[A]] interface {
	~func(M, B) A
}

type Functor[A, B, HKTA, HKTB any] interface {
	~func(HKTA, func(A) B) HKTB
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
