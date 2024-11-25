package repository

import (
	"context"
	"io/fs"
	"syscall"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
	"github.com/unmango/go/fs/github/repository/release"
)

type File struct {
	internal.ReadOnlyFile
	client *github.Client
	owner  string
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
	return release.Readdir(context.TODO(),
		f.client,
		f.owner,
		f.repo.GetName(),
		count,
	)
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	return release.Readdirnames(context.TODO(),
		f.client,
		f.owner,
		f.repo.GetName(),
		n,
	)
}

// Seek implements afero.File.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (f *File) Stat() (fs.FileInfo, error) {
	return &FileInfo{repo: f.repo}, nil
}

func NewFile(gh *github.Client, owner string, repository *github.Repository) *File {
	return &File{
		client: gh,
		owner:  owner,
		repo:   repository,
	}
}
