package either

type Either[L, R any] func() (L, R)

func Left[L, R any](left L) Either[L, R] {
	return From(left, *new(R))
}

func Right[L, R any](right R) Either[L, R] {
	return From(*new(L), right)
}

func From[L, R any](left L, right R) Either[L, R] {
	return func() (L, R) {
		return left, right
	}
}

func Map[LA, LB, RA, RB comparable, E Either[LA, RA]](either E, fn func(LA, RA) (LB, RB)) Either[LB, RB] {
	return func() (LB, RB) {
		return fn(either())
	}
}

func MapLeft[LA, LB, R comparable, E Either[LA, R]](either E, fn func(LA) LB) Either[LB, R] {
	return Map(either, func(l LA, r R) (LB, R) {
		if isZero(l) {
			return *new(LB), r
		} else {
			return fn(l), r
		}
	})
}

func MapRight[L, RA, RB comparable, E Either[L, RA]](either E, fn func(RA) RB) Either[L, RB] {
	return Map(either, func(l L, r RA) (L, RB) {
		if isZero(r) {
			return l, *new(RB)
		} else {
			return l, fn(r)
		}
	})
}

func Bind[LA, LB, RA, RB comparable, E Either[LA, RA]](either E, fn func(LA, RA) Either[LB, RB]) Either[LB, RB] {
	return func() (LB, RB) {
		return fn(either())()
	}
}

func BindLeft[LA, LB, R comparable, A Either[LA, R]](either A, fn func(LA) Either[LB, R]) Either[LB, R] {
	return Bind(either, func(l LA, r R) Either[LB, R] {
		if isZero(l) {
			return Right[LB](r)
		} else {
			return fn(l)
		}
	})
}

func BindRight[L, RA, RB comparable, E Either[L, RA]](either E, fn func(RA) Either[L, RB]) Either[L, RB] {
	return Bind(either, func(l L, r RA) Either[L, RB] {
		if isZero(r) {
			return Left[L, RB](l)
		} else {
			return fn(r)
		}
	})
}

func isZero[T comparable](t T) bool {
	return t == *new(T)
}
