package user

import (
	"context"
	"io/fs"
	"syscall"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
	L "github.com/unmango/go/lazy"
)

type UserResult = internal.Result[*github.User]

type File struct {
	internal.ReadOnlyFile
	L.Lazy[UserResult]
	name string
}

// Close implements afero.File.
func (f *File) Close() error {
	_, res, err := f.value()
	if err != nil {
		return err
	}

	return res.Body.Close()
}

// Name implements afero.File.
func (f *File) Name() string {
	return f.name
}

// Read implements afero.File.
func (f *File) Read(p []byte) (n int, err error) {
	_, res, err := f.value()
	if err != nil {
		return 0, err
	}

	return res.Body.Read(p)
}

// ReadAt implements afero.File.
func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (f *File) Readdir(count int) ([]fs.FileInfo, error) {
	panic("unimplemented")
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	panic("unimplemented")
}

// Seek implements afero.File.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
// Subtle: this method shadows the method (*Fs).Stat of user.Fs.
func (f *File) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}

func (f *File) value() (*github.User, *github.Response, error) {
	return f.Lazy()()
}

func NewFile(gh *github.Client, name string) *File {
	get := func() (*github.User, *github.Response, error) {
		return gh.Users.Get(context.TODO(), name)
	}

	return &File{
		name: name,
		Lazy: internal.Request(get),
	}
}
