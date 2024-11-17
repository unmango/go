package writer

import (
	"io"
	"io/fs"
	"syscall"
	"time"

	"github.com/spf13/afero"
)

type Fs struct{ io.Writer }

// Chmod implements afero.Fs.
func (w *Fs) Chmod(string, fs.FileMode) error {
	return syscall.EPERM
}

// Chown implements afero.Fs.
func (w *Fs) Chown(string, int, int) error {
	return syscall.EPERM
}

// Chtimes implements afero.Fs.
func (w *Fs) Chtimes(string, time.Time, time.Time) error {
	return syscall.EPERM
}

// Create implements afero.Fs.
func (w *Fs) Create(name string) (afero.File, error) {
	return &File{w.Writer, name}, nil
}

// Mkdir implements afero.Fs.
func (w *Fs) Mkdir(name string, perm fs.FileMode) error {
	return syscall.EPERM
}

// MkdirAll implements afero.Fs.
func (w *Fs) MkdirAll(path string, perm fs.FileMode) error {
	return syscall.EPERM
}

// Name implements afero.Fs.
func (w *Fs) Name() string {
	return "io.Writer"
}

// Open implements afero.Fs.
func (w *Fs) Open(name string) (afero.File, error) {
	return &File{w.Writer, name}, nil
}

// OpenFile implements afero.Fs.
func (w *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return &File{w.Writer, name}, nil
}

// Remove implements afero.Fs.
func (w *Fs) Remove(name string) error {
	return syscall.EPERM
}

// RemoveAll implements afero.Fs.
func (w *Fs) RemoveAll(path string) error {
	return syscall.EPERM
}

// Rename implements afero.Fs.
func (w *Fs) Rename(oldname string, newname string) error {
	return syscall.EPERM
}

// Stat implements afero.Fs.
func (w *Fs) Stat(name string) (fs.FileInfo, error) {
	return &FileInfo{w.Writer, name}, nil
}
