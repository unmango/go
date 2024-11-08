package repository

import (
	"io/fs"

	"github.com/google/go-github/v66/github"
)

type fileWriter struct {
	name   string
	client *github.RepositoriesService
}

// Close implements afero.File.
func (f *fileWriter) Close() error {
	panic("unimplemented")
}

// Name implements afero.File.
func (f *fileWriter) Name() string {
	panic("unimplemented")
}

// Read implements afero.File.
func (f *fileWriter) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// ReadAt implements afero.File.
func (f *fileWriter) ReadAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// Readdir implements afero.File.
func (f *fileWriter) Readdir(count int) ([]fs.FileInfo, error) {
	panic("unimplemented")
}

// Readdirnames implements afero.File.
func (f *fileWriter) Readdirnames(n int) ([]string, error) {
	panic("unimplemented")
}

// Seek implements afero.File.
func (f *fileWriter) Seek(offset int64, whence int) (int64, error) {
	panic("unimplemented")
}

// Stat implements afero.File.
func (f *fileWriter) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}

// Sync implements afero.File.
func (f *fileWriter) Sync() error {
	panic("unimplemented")
}

// Truncate implements afero.File.
func (f *fileWriter) Truncate(size int64) error {
	panic("unimplemented")
}

// Write implements afero.File.
func (f *fileWriter) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}

// WriteAt implements afero.File.
func (f *fileWriter) WriteAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// WriteString implements afero.File.
func (f *fileWriter) WriteString(s string) (ret int, err error) {
	panic("unimplemented")
}
