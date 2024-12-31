package version

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/unmango/go/devops/work"
	"github.com/unmango/go/option"
)

type Options struct {
	fs   afero.Fs
	root string
}

type Option func(*Options)

func RelPath(name string) string {
	return filepath.Join(".versions", name)
}

func ReadFile(name string, options ...Option) (string, error) {
	opts := &Options{fs: afero.NewOsFs()}
	option.ApplyAll(opts, options)

	p := filepath.Join(opts.root, RelPath(name))
	if b, err := afero.ReadFile(opts.fs, p); err == nil {
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

func WithRoot(root string) Option {
	return func(o *Options) {
		o.root = root
	}
}

func WithWorkspace(work *work.Context) Option {
	return WithRoot(work.Root())
}
