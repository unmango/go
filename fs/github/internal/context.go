package internal

import "context"

type ctxval struct{ ctx context.Context }

// Context implements ContextAccessor.
func (b *ctxval) Context() context.Context {
	return b.ctx
}

type ContextAccessor interface {
	Context() context.Context
}

func BackgroundContext() ContextAccessor {
	return &ctxval{context.Background()}
}
