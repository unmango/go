package signal

// Map implements Functor for Signal[T]. It creates a new
// signal, and therefore does not notify subscribers
func Map[T, V any](s Signal[T], f func(T) V) Signal[V] {
	return New(f(s.Get()))
}
