package signal

import "github.com/unmango/go/rx"

// Map implements Functor for Signal[T]. It creates a new
// signal, and therefore does not notify subscribers
func Map[T, V any](s rx.Signal[T], f func(T) V) rx.Signal[V] {
	return New(f(s.Get()))
}
