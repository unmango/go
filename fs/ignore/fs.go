package ignore

import (
	"fmt"
	"io"

	"github.com/ianlewis/go-gitignore"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/filter"
)

func NewGitFs(base afero.Fs, ignore gitignore.GitIgnore) afero.Fs {
	return filter.NewFs(base, ignore.Ignore)
}

func FromGitIgnore(base afero.Fs, reader io.Reader) (afero.Fs, error) {
	if i, err := readAll(reader); err != nil {
		return nil, fmt.Errorf("reading ignore file: %w", err)
	} else {
		return NewGitFs(base, i), nil
	}
}

func FromGitIgnoreFile(base afero.Fs, path string) (afero.Fs, error) {
	f, err := base.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening ignore file: %w", err)
	}
	defer f.Close()

	return FromGitIgnore(base, f)
}

func LoadDefaultGitIgnore(base afero.Fs) (afero.Fs, error) {
	return FromGitIgnoreFile(base, ".gitignore")
}

func readAll(reader io.Reader) (gitignore.GitIgnore, error) {
	var err error
	i := gitignore.New(reader, "",
		func(e gitignore.Error) bool {
			err = e
			return false
		},
	)
	if err != nil {
		return nil, err
	} else {
		return i, nil
	}
}
