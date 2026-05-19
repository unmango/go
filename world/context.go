package world

import "context"

type key struct{}

func FromContext(ctx context.Context) World {
	if w, ok := ctx.Value(key{}).(World); ok {
		return w
	}
	return System
}

func WithContext(parent context.Context, w World) context.Context {
	return context.WithValue(parent, key{}, w)
}
