package user

import (
	"io/fs"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
)

type File struct {
	internal.ReadOnlyFile
	internal.Fs
	User *github.User
}

// Close implements afero.File.
func (File) Close() error {
	return nil
}

// Name implements afero.File.
func (u *File) Name() string {
	return u.User.GetName()
}

// Read implements afero.File.
func (u *File) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// ReadAt implements afero.File.
func (u *File) ReadAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// Readdir implements afero.File.
func (u *File) Readdir(count int) ([]fs.FileInfo, error) {
	panic("unimplemented")
}

// Readdirnames implements afero.File.
func (u *File) Readdirnames(n int) ([]string, error) {
	panic("unimplemented")
}

// Seek implements afero.File.
func (u *File) Seek(offset int64, whence int) (int64, error) {
	panic("unimplemented")
}

// Stat implements afero.File.
// Subtle: this method shadows the method (*Fs).Stat of user.Fs.
func (u *File) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}
