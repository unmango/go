package user

import (
	"io/fs"
	"path/filepath"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
)

type userFile struct {
	internal.ReadOnlyFile
	user *github.User

	repos map[int64]*github.Repository
}

// Close implements afero.File.
func (userFile) Close() error {
	return nil
}

// Name implements afero.File.
// Subtle: this method shadows the method (*Fs).Name of user.Fs.
func (u *userFile) Name() string {
	return filepath.Join(u.parent.Name(), u.user.GetName())
}

// Read implements afero.File.
func (u *userFile) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// ReadAt implements afero.File.
func (u *userFile) ReadAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// Readdir implements afero.File.
func (u *userFile) Readdir(count int) ([]fs.FileInfo, error) {
	panic("unimplemented")
}

// Readdirnames implements afero.File.
func (u *userFile) Readdirnames(n int) ([]string, error) {
	panic("unimplemented")
}

// Seek implements afero.File.
func (u *userFile) Seek(offset int64, whence int) (int64, error) {
	panic("unimplemented")
}

// Stat implements afero.File.
// Subtle: this method shadows the method (*Fs).Stat of user.Fs.
func (u *userFile) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}
