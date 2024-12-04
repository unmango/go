package filter

import (
	"fmt"
	"io/fs"
	"syscall"
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
	if f.pred != nil && f.pred(name) {
		return f.src.Chmod(name, mode)
	}

	return syscall.ENOENT
}

// Chown implements afero.Fs.
func (f *Fs) Chown(name string, uid int, gid int) error {
	if f.pred != nil && f.pred(name) {
		return f.src.Chown(name, uid, gid)
	}

	return syscall.ENOENT
}

// Chtimes implements afero.Fs.
func (f *Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	if f.pred != nil && f.pred(name) {
		return f.src.Chtimes(name, atime, mtime)
	}

	return syscall.ENOENT
}

// Create implements afero.Fs.
func (f *Fs) Create(name string) (afero.File, error) {
	if f.pred != nil && f.pred(name) {
		return f.src.Create(name)
	}

	return nil, syscall.ENOENT
}

// Mkdir implements afero.Fs.
func (f *Fs) Mkdir(name string, perm fs.FileMode) error {
	return f.src.Mkdir(name, perm)
}

// MkdirAll implements afero.Fs.
func (f *Fs) MkdirAll(path string, perm fs.FileMode) error {
	return f.src.Mkdir(path, perm)
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("Filter: %s", f.src.Name())
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	if f.pred != nil && f.pred(name) {
		return f.src.Open(name)
	}

	return nil, syscall.ENOENT
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	if f.pred != nil && f.pred(name) {
		return f.src.OpenFile(name, flag, perm)
	}

	return nil, syscall.ENOENT
}

// Remove implements afero.Fs.
func (f *Fs) Remove(name string) error {
	stat, err := f.src.Stat(name)
	if err != nil {
		return err
	}
	if stat.IsDir() || (f.pred != nil && f.pred(name)) {
		return f.src.Remove(name)
	}

	return syscall.ENOENT
}

// RemoveAll implements afero.Fs.
func (f *Fs) RemoveAll(path string) error {
	if f.pred != nil && f.pred(path) {
		return f.src.Remove(path)
	}

	return syscall.ENOENT
}

// Rename implements afero.Fs.
func (f *Fs) Rename(oldname string, newname string) error {
	if f.pred != nil && f.pred(oldname) {
		return f.src.Rename(oldname, newname)
	}

	return syscall.ENOENT
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	if f.pred != nil && f.pred(name) {
		return f.src.Stat(name)
	}

	return nil, syscall.ENOENT
}

func NewFs(src afero.Fs, predicate Predicate) afero.Fs {
	return &Fs{src: src, pred: predicate}
}
