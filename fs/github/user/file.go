package user

import (
	"context"
	"io/fs"
	"syscall"

	"github.com/google/go-github/v67/github"
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
	return f.user.GetLogin()
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
	return repository.Readdir(context.TODO(), f.client, f.Name(), count)
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	return repository.Readdirnames(context.TODO(), f.client, f.Name(), n)
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

func NewFile(gh *github.Client, user *github.User) *File {
	return &File{
		client: gh,
		user:   user,
	}
}
