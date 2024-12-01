package content

import (
	"io/fs"
	"syscall"

	"github.com/google/go-github/v67/github"
	"github.com/unmango/go/fs/github/internal"
)

type Directory struct {
	internal.ReadOnlyFile
	internal.ContentPath

	client  *github.Client
	content []*github.RepositoryContent
}

// Close implements afero.File.
func (d *Directory) Close() error {
	return nil
}

// Name implements afero.File.
func (d *Directory) Name() string {
	return d.Content
}

// Read implements afero.File.
func (d *Directory) Read(p []byte) (n int, err error) {
	return 0, syscall.EPERM
}

// ReadAt implements afero.File.
func (d *Directory) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (d *Directory) Readdir(count int) ([]fs.FileInfo, error) {
	length := min(count, len(d.content))
	files := make([]fs.FileInfo, length)

	for i, c := range d.content {
		files[i] = &FileInfo{content: c}
	}

	return files, nil
}

// Readdirnames implements afero.File.
func (d *Directory) Readdirnames(n int) ([]string, error) {
	length := min(n, len(d.content))
	names := make([]string, length)

	for i, c := range d.content {
		names[i] = c.GetName()
	}

	return names, nil
}

// Seek implements afero.File.
func (d *Directory) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (d *Directory) Stat() (fs.FileInfo, error) {
	return &DirectoryInfo{
		name:    d.Name(),
		content: d.content,
	}, nil
}
