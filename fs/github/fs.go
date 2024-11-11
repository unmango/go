package github

import (
	"io/fs"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/internal"
	"github.com/unmango/go/fs/github/user"
	"github.com/unmango/go/option"
)

type Fs struct {
	afero.ReadOnlyFs
	internal.Fs
}

var _ afero.Fs = &Fs{}

// Name implements afero.Fs.
func (g *Fs) Name() string {
	return "https://github.com"
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	u, err := f.getUser(name)
	if err != nil {
		return nil, err
	}

	return &user.File{
		Fs:   f.Fs,
		User: u,
	}, nil
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return f.Open(name)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	u, err := f.getUser(name)
	if err != nil {
		return nil, err
	}

	return &user.FileInfo{
		Fs:   f.Fs,
		User: u,
	}, nil
}

func (f *Fs) getUser(name string) (*github.User, error) {
	user, _, err := f.Client.Users.Get(f.Context(), name)
	return user, err
}

func New(owner, repository string, options ...Option) *Fs {
	gh := &Fs{}
	option.ApplyAll(gh, options)

	if gh.Client == nil {
		gh.Client = DefaultClient()
	}

	return gh
}
