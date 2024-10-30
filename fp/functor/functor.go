package functor

import "github.com/unmango/go/fp/constraint"

// Still not sure any of this is right, but I'm learning

type Functor[
	A, B any,
	HKTA constraint.Type[A],
	HKTB constraint.Type[B],
] interface {
	Map(HKTA, func(A) B) HKTB
}

type functor[
	A, B any,
	HKTA constraint.Type[A],
	HKTB constraint.Type[B],
] func(HKTA, func(A) B) HKTB

func (f functor[A, B, HKTA, HKTB]) Map(a HKTA, fn func(A) B) HKTB {
	return f(a, fn)
}

func Lift[
	A, B any,
	HKTA constraint.Type[A],
	HKTB constraint.Type[B],
	F constraint.Functor[A, B, HKTA, HKTB],
](f F) Functor[A, B, HKTA, HKTB] {
	return functor[A, B, HKTA, HKTB](f)
}
