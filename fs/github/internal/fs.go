package internal

import "github.com/spf13/afero"

type Fs struct {
	afero.Fs
	*State
}
