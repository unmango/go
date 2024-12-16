package aferox

import (
	"context"
	"fmt"

	"github.com/spf13/afero"
)

type ContextSetter interface {
	SetContext(context.Context)
}

type ContextAccessor interface {
	Context() context.Context
}

func SetContext(fs afero.Fs, ctx context.Context) error {
	if ctxfs, ok := fs.(ContextSetter); !ok {
		return fmt.Errorf("context not supported: %s", fs.Name())
	} else {
		ctxfs.SetContext(ctx)
	}

	return nil
}
