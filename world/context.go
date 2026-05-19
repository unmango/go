package world

import "context"

type key struct{}

func FromContext(ctx context.Context) IO {
	if w, ok := ctx.Value(key{}).(IO); ok {
		return w
	}
	return System
}

func WithContext(parent context.Context, w IO) context.Context {
	return context.WithValue(parent, key{}, w)
}
