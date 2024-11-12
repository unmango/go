package user

import (
	"context"
	"fmt"
	"io/fs"
	"syscall"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
	"github.com/unmango/go/fs/github/repository"
)

type File struct {
	internal.ReadOnlyFile
	client *github.Client
	user   *github.User
}

// Close implements afero.File.
func (f *File) Close() error {
	return nil
}

// Name implements afero.File.
func (f *File) Name() string {
	return f.user.GetName()
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
	files, err := repository.Readdir(context.TODO(), f.client, f.user.GetName(), count)
	if err != nil {
		return nil, fmt.Errorf("readdir: %w", err)
	}

	results := make([]fs.FileInfo, len(files))
	for i, f := range files {
		results[i] = f
	}

	return results, nil
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	files, err := repository.Readdir(context.TODO(), f.client, f.user.GetName(), n)
	if err != nil {
		return nil, fmt.Errorf("readdirnames: %w", err)
	}

	names := make([]string, len(files))
	for i, d := range files {
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
		user:   f.user,
	}, nil
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
