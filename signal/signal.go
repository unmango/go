package signal

type Subscriber[T any] func(T)

type Signal[T any] interface {
	Get() T
	Set(T)
	Subscribe(Subscriber[T])
}

type signal[T any] struct {
	value       T
	subscribers []Subscriber[T]
}

func (s signal[T]) Get() T {
	return s.value
}

func (s *signal[T]) Set(v T) {
	s.value = v
}

func (s *signal[T]) notify() {
	for _, sub := range s.subscribers {
		sub(s.value)
	}
}

func (s *signal[T]) Subscribe(sub Subscriber[T]) {
	s.subscribers = append(s.subscribers, sub)
}

func New[T any](v T) Signal[T] {
	return &signal[T]{v, []Subscriber[T]{}}
}
