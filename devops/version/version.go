package version

import (
	"path/filepath"

	"github.com/spf13/afero"
)

const DirName = ".versions"

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

func RelPath(name string) string {
	return filepath.Join(DirName, name)
}

func WithFs(fs afero.Fs) Option {
	return func(o *Options) {
		o.fs = fs
	}
}
