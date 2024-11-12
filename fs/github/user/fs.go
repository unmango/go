package user

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/repository"
)

type Fs struct {
	afero.ReadOnlyFs
	client *github.Client
	user   string
}

// Name implements afero.Fs.
func (g *Fs) Name() string {
	return fmt.Sprintf("https://github.com/%s", g.user)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	return repository.Open(context.TODO(), f.client, f.user, name)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return repository.Open(context.TODO(), f.client, f.user, name)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	return repository.Stat(context.TODO(), f.client, f.user, name)
}

func New(gh *github.Client, name string) afero.Fs {
	return &Fs{
		client: gh,
		user:   name,
	}
}
