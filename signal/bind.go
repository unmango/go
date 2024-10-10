package signal

func Bind[T, V any](s Signal[T], f func(T) Signal[V]) Signal[V] {
	return f(s.Get())
}
