package maybe

type Maybe[T any] func() (T, error)

type e struct{}

// Error implements error.
func (e) Error() string { return "None" }

var (
	ErrNone error = e{}
	ptrs    map[interface{}]interface{}
)

func Ok[T any](v T) Maybe[T] {
	return func() (T, error) {
		return v, nil
	}
}

func None[T any]() (T, error) {
	return zero[T](), ErrNone
}

func IsOk[T any](m Maybe[T]) bool {
	return !IsNone(m)
}

func IsNone[T any](m Maybe[T]) bool {
	n := None[T]
	if p, ok := ptrs[&n]; ok {
		return true
	} else {
		
	}
}

func zero[T any]() T {
	// TODO: Something something allocations
	var t T
	return t
}
