package aferox

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/unmango/go/fs/context"
)

func SetContext(fs afero.Fs, ctx context.Context) error {
	if ctxfs, ok := fs.(context.Setter); !ok {
		return fmt.Errorf("context not supported: %s", fs.Name())
	} else {
		ctxfs.SetContext(ctx)
	}

	return nil
}
