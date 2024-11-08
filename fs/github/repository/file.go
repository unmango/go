package repository

import (
	"bytes"
	"errors"
	"io/fs"

	"github.com/google/go-github/v66/github"
)

type fileContent struct {
	file *github.RepositoryContent
	buf  *bytes.Buffer
}

// Close implements afero.File.
func (r *fileContent) Close() error { return nil }

// Name implements afero.File.
func (r *fileContent) Name() string { return r.file.GetName() }

// Read implements afero.File.
func (r *fileContent) Read(p []byte) (n int, err error) {
	if err = r.ensureBuffer(); err != nil {
		return
	}

	return r.buf.Read(p)
}

// ReadAt implements afero.File.
func (r *fileContent) ReadAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// Readdir implements afero.File.
func (r *fileContent) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, errors.New("not a directory")
}

// Readdirnames implements afero.File.
func (r *fileContent) Readdirnames(n int) ([]string, error) {
	return nil, errors.New("not a directory")
}

// Seek implements afero.File.
func (r *fileContent) Seek(offset int64, whence int) (int64, error) {
	panic("unimplemented")
}

// Stat implements afero.File.
func (r *fileContent) Stat() (fs.FileInfo, error) {
	return &fileInfoContent{r.file}, nil
}

// Sync implements afero.File.
func (r *fileContent) Sync() error {
	panic("unimplemented")
}

// Truncate implements afero.File.
func (r *fileContent) Truncate(size int64) error {
	panic("unimplemented")
}

// Write implements afero.File.
func (r *fileContent) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}

// WriteAt implements afero.File.
func (r *fileContent) WriteAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// WriteString implements afero.File.
func (r *fileContent) WriteString(s string) (ret int, err error) {
	panic("unimplemented")
}

func (r *fileContent) ensureBuffer() error {
	c, err := r.file.GetContent()
	if err != nil {
		return err
	}

	r.buf = bytes.NewBufferString(c)
	return nil
}

func NewFile(contents *github.RepositoryContent) fs.File {
	return &fileContent{file: contents}
}
