package aferox

import (
	"io/fs"

	"github.com/spf13/afero"
)

type ReduceFunc[T any] func(string, fs.FileInfo, T, error) (T, error)

func Reduce[T any](fsys afero.Fs, root string, reduce ReduceFunc[T], initial T) (acc T, err error) {
	acc = initial
	err = afero.Walk(fsys, root,
		func(path string, info fs.FileInfo, err error) error {
			acc, err = reduce(path, info, acc, err)
			return err
		},
	)

	return
}
