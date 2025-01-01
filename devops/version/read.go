package version

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/unmango/go/option"
)

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

func RelPath(name string) string {
	return filepath.Join(DirName, name)
}

func ReadFile(name string, options ...Option) (string, error) {
	opts := &Options{fs: afero.NewOsFs()}
	option.ApplyAll(opts, options)

	if b, err := afero.ReadFile(opts.fs, RelPath(name)); err == nil {
		return strings.TrimSpace(string(b)), nil
	} else if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("dependency not found: %s", name)
	} else {
		return "", err
	}
}

func WithFs(fs afero.Fs) Option {
	return func(o *Options) {
		o.fs = fs
	}
}
