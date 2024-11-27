package content

import (
	"io/fs"
	"syscall"

	"github.com/google/go-github/v67/github"
	"github.com/unmango/go/fs/github/internal"
)

type Directory struct {
	internal.ReadOnlyFile
	client   *github.Client
	name     string
	contents []*github.RepositoryContent
}

// Close implements afero.File.
func (d *Directory) Close() error {
	return nil
}

// Name implements afero.File.
func (d *Directory) Name() string {
	return d.name
}

// Read implements afero.File.
func (d *Directory) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// ReadAt implements afero.File.
func (d *Directory) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (d *Directory) Readdir(count int) ([]fs.FileInfo, error) {
	length := min(count, len(d.contents))
	files := make([]fs.FileInfo, length)

	for i, c := range d.contents {
		files[i] = &FileInfo{content: c}
	}

	return files, nil
}

// Readdirnames implements afero.File.
func (d *Directory) Readdirnames(n int) ([]string, error) {
	length := min(n, len(d.contents))
	names := make([]string, length)

	for i, c := range d.contents {
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
		name:     d.name,
		contents: d.contents,
	}, nil
}
