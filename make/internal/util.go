package internal

// https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang

type Stack[T any] []T

func (s Stack[T]) Push(v T) Stack[T] {
	return append(s, v)
}

func (s Stack[T]) Pop() (Stack[T], T, bool) {
	if len(s) == 0 {
		var t T
		return nil, t, false
	}

	l := len(s)
	return s[:l-1], s[l-1], true
}
