package repository

import (
	"context"
	"fmt"
	"io/fs"
	"syscall"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
)

type File struct {
	internal.ReadOnlyFile
	client *github.Client
	repo   *github.Repository
}

// Close implements afero.File.
func (f *File) Close() error {
	return nil
}

// Name implements afero.File.
func (f *File) Name() string {
	return f.repo.GetName()
}

// Read implements afero.File.
func (f *File) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// ReadAt implements afero.File.
func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (f *File) Readdir(count int) ([]fs.FileInfo, error) {
	// return repository.Readdir(context.TODO(), f.client, f.user.GetName(), count)
	panic("unimplemented")
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	dirs, err := f.Readdir(n)
	if err != nil {
		return nil, fmt.Errorf("readdirname: %w", err)
	}

	names := make([]string, len(dirs))
	for i, d := range dirs {
		names[i] = d.Name()
	}

	return names, nil
}

// Seek implements afero.File.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (f *File) Stat() (fs.FileInfo, error) {
	return &FileInfo{
		client: f.client,
		repo:   f.repo,
	}, nil
}

func Open(ctx context.Context, gh *github.Client, user, name string) (*File, error) {
	repo, _, err := gh.Repositories.Get(ctx, user, name)
	if err != nil {
		return nil, err
	}

	return &File{
		client: gh,
		repo:   repo,
	}, nil
}
