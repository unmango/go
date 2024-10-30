package pipe

import "github.com/unmango/go/fp/constraint"

func Pipe[
	A, B, C any,
	FA constraint.Map[A, B],
	FB constraint.Map[B, C],
	R constraint.Map[A, C],
](fa FA, fb FB) R {
	return func(a A) C {
		return fb(fa(a))
	}
}

func Pipe3[
	A, B, C, D any,
	FA constraint.Map[A, B],
	FB constraint.Map[B, C],
	FC constraint.Map[C, D],
	R constraint.Map[A, D],
](fa FA, fb FB, fc FC) R {
	return func(a A) D {
		return fc(fb(fa(a)))
	}
}

func Pipe4[
	A, B, C, D, E any,
	FA constraint.Map[A, B],
	FB constraint.Map[B, C],
	FC constraint.Map[C, D],
	FD constraint.Map[D, E],
	R constraint.Map[A, E],
](fa FA, fb FB, fc FC, fd FD) R {
	return func(a A) E {
		return fd(fc(fb(fa(a))))
	}
}

func Pipe5[
	A, B, C, D, E, F any,
	FA constraint.Map[A, B],
	FB constraint.Map[B, C],
	FC constraint.Map[C, D],
	FD constraint.Map[D, E],
	FE constraint.Map[E, F],
	R constraint.Map[A, F],
](fa FA, fb FB, fc FC, fd FD, fe FE) R {
	return func(a A) F {
		return fe(fd(fc(fb(fa(a)))))
	}
}

func Pipe6[
	A, B, C, D, E, F, G any,
	FA constraint.Map[A, B],
	FB constraint.Map[B, C],
	FC constraint.Map[C, D],
	FD constraint.Map[D, E],
	FE constraint.Map[E, F],
	FF constraint.Map[F, G],
	R constraint.Map[A, G],
](fa FA, fb FB, fc FC, fd FD, fe FE, ff FF) R {
	return func(a A) G {
		return ff(fe(fd(fc(fb(fa(a))))))
	}
}
