package functor

import "github.com/unmango/go/fp/constraint"

// Still not sure any of this is right, but I'm learning

type Functor[
	A, B any,
	HKTA constraint.Type[A],
	HKTB constraint.Type[B],
	M constraint.Map[A, B],
] interface {
	Map(HKTA, M) HKTB
}

type functor[
	A, B any,
	HKTA constraint.Type[A],
	HKTB constraint.Type[B],
	M constraint.Map[A, B],
] func(HKTA, M) HKTB

func (f functor[A, B, HKTA, HKTB, M]) Map(a HKTA, fn M) HKTB {
	return f(a, fn)
}

func Lift[
	A, B any,
	HKTA constraint.Type[A],
	HKTB constraint.Type[B],
	M constraint.Map[A, B],
	F constraint.Functor[A, B, HKTA, HKTB, M],
](f F) Functor[A, B, HKTA, HKTB, M] {
	return functor[A, B, HKTA, HKTB, M](f)
}
