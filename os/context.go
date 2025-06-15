package os

import "context"

type key struct{}

func FromContext(ctx context.Context) Os {
	if v := ctx.Value(key{}); v != nil {
		return v.(Os)
	} else {
		return System
	}
}

func WithContext(parent context.Context, val Os) context.Context {
	return context.WithValue(parent, key{}, val)
}
