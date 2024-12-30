package version

import (
	"context"
	"path/filepath"

	"github.com/spf13/afero"
	aferox "github.com/unmango/go/fs"
)

func Sprint(ctx context.Context, name string) (string, error) {
	fs := aferox.FromContext(ctx)
	p := filepath.Join(".versions", name)

	if b, err := afero.ReadFile(fs, p); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}
