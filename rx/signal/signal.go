package signal

import "github.com/unmango/go/rx"

type Subscriber[T any] func(T)

type signal[T any] struct {
	value T
	subs  []Subscriber[T]
}

func (s signal[T]) Get() T {
	return s.value
}

func (s *signal[T]) Set(v T) {
	s.value = v
	s.notify()
}

func (s *signal[T]) notify() {
	for _, sub := range s.subs {
		sub(s.value)
	}
}

func (s *signal[T]) Subscribe(sub func(T)) {
	s.subs = append(s.subs, sub)
}

func New[T any](v T) rx.Signal[T] {
	return &signal[T]{v, []Subscriber[T]{}}
}
