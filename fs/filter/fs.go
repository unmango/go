package filter

import (
	"io/fs"
	"time"

	"github.com/spf13/afero"
)

type Predicate func(string) bool

type Fs struct {
	src  afero.Fs
	pred Predicate
}

// Chmod implements afero.Fs.
func (f *Fs) Chmod(name string, mode fs.FileMode) error {
	panic("unimplemented")
}

// Chown implements afero.Fs.
func (f *Fs) Chown(name string, uid int, gid int) error {
	panic("unimplemented")
}

// Chtimes implements afero.Fs.
func (f *Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	panic("unimplemented")
}

// Create implements afero.Fs.
func (f *Fs) Create(name string) (afero.File, error) {
	panic("unimplemented")
}

// Mkdir implements afero.Fs.
func (f *Fs) Mkdir(name string, perm fs.FileMode) error {
	panic("unimplemented")
}

// MkdirAll implements afero.Fs.
func (f *Fs) MkdirAll(path string, perm fs.FileMode) error {
	panic("unimplemented")
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	panic("unimplemented")
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	panic("unimplemented")
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	panic("unimplemented")
}

// Remove implements afero.Fs.
func (f *Fs) Remove(name string) error {
	panic("unimplemented")
}

// RemoveAll implements afero.Fs.
func (f *Fs) RemoveAll(path string) error {
	panic("unimplemented")
}

// Rename implements afero.Fs.
func (f *Fs) Rename(oldname string, newname string) error {
	panic("unimplemented")
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	panic("unimplemented")
}

func NewFs(src afero.Fs, predicate Predicate) afero.Fs {
	return &Fs{src: src, pred: predicate}
}
