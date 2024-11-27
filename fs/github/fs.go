package github

import (
	"context"
	"io/fs"

	"github.com/google/go-github/v67/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/internal"
	"github.com/unmango/go/fs/github/user"
)

type Fs struct {
	afero.ReadOnlyFs
	client *github.Client
}

// Name implements afero.Fs.
func (g *Fs) Name() string {
	return "https://github.com"
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	return user.Open(context.TODO(), f.client, name)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return user.Open(context.TODO(), f.client, name)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	return user.Stat(context.TODO(), f.client, name)
}

func NewFs(gh *github.Client) afero.Fs {
	if gh == nil {
		gh = internal.DefaultClient()
	}

	return &Fs{client: gh}
}
