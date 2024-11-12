package user

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
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
	return Open(context.TODO(), f.client, name)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return Open(context.TODO(), f.client, name)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	return Stat(context.TODO(), f.client, name)
}

func NewFs(gh *github.Client) afero.Fs {
	return &Fs{client: gh}
}

func Open(ctx context.Context, gh *github.Client, name string) (*File, error) {
	user, _, err := gh.Users.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("open user: %w", err)
	}

	return &File{
		client: gh,
		user:   user,
	}, nil
}

func Stat(ctx context.Context, gh *github.Client, name string) (*FileInfo, error) {
	user, _, err := gh.Users.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("stat user: %w", err)
	}

	return &FileInfo{
		client: gh,
		user:   user,
	}, nil
}
