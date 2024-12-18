package aferox

import (
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/internal"
)

func Copy(src, dest afero.Fs) error {
	return internal.Copy(src, dest)
}
