package either

type Either[L, R any] func() (*L, *R)

func Left[R, L any](left L) Either[L, R] {
	return func() (*L, *R) {
		return &left, nil
	}
}

func Right[L, R any](right R) Either[L, R] {
	return func() (*L, *R) {
		return nil, &right
	}
}

func MapLeft[LA, LB, R any, E Either[LA, R]](either E, fn func(LA) LB) Either[LB, R] {
	return func() (*LB, *R) {
		if l, r := either(); l != nil {
			lb := fn(*l)
			return &lb, nil
		} else {
			return nil, r
		}
	}
}

func MapRight[L, RA, RB any, E Either[L, RA]](either E, fn func(RA) RB) Either[L, RB] {
	return func() (*L, *RB) {
		if l, r := either(); r != nil {
			rb := fn(*r)
			return nil, &rb
		} else {
			return l, nil
		}
	}
}

func BindLeft[LA, LB, R any, A Either[LA, R]](either A, fn func(LA) Either[LB, R]) Either[LB, R] {
	return func() (*LB, *R) {
		if l, r := either(); l != nil {
			eb := fn(*l)
			return eb()
		} else {
			return nil, r
		}
	}
}

func BindRight[L, RA, RB any, E Either[L, RA]](either E, fn func(RA) Either[L, RB]) Either[L, RB] {
	return func() (*L, *RB) {
		if l, r := either(); r != nil {
			eb := fn(*r)
			return eb()
		} else {
			return l, nil
		}
	}
}
