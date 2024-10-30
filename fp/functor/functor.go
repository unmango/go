package functor

import "github.com/unmango/go/fp/constraints"

type Functor[A, B, HKTA, HKTB any] interface {
	Map(HKTA, func(A) B) HKTB
}

type functor[A, B, HKTA, HKTB any] func(HKTA, func(A) B) HKTB

func (f functor[A, B, HKTA, HKTB]) Map(a HKTA, fn func(A) B) HKTB {
	return f(a, fn)
}

func Lift[
	A, B, HKTA, HKTB any,
	F constraints.Functor[A, B, HKTA, HKTB],
](f F) Functor[A, B, HKTA, HKTB] {
	return functor[A, B, HKTA, HKTB](f)
}
