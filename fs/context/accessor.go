package context

import (
	"fmt"
	"io/fs"
	"time"

	"github.com/spf13/afero"
	aferox "github.com/unmango/go/fs"
)

type AccessorFunc func() Context

func (fn AccessorFunc) Context() Context {
	return fn()
}

type AccessorFs struct {
	aferox.ContextAccessor
	fs Fs
}

// Chmod implements afero.Fs.
func (a *AccessorFs) Chmod(name string, mode fs.FileMode) error {
	return a.fs.Chmod(a.Context(), name, mode)
}

// Chown implements afero.Fs.
func (a *AccessorFs) Chown(name string, uid int, gid int) error {
	return a.fs.Chown(a.Context(), name, uid, gid)
}

// Chtimes implements afero.Fs.
func (a *AccessorFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return a.fs.Chtimes(a.Context(), name, atime, mtime)
}

// Create implements afero.Fs.
func (a *AccessorFs) Create(name string) (afero.File, error) {
	return a.fs.Create(a.Context(), name)
}

// Mkdir implements afero.Fs.
func (a *AccessorFs) Mkdir(name string, perm fs.FileMode) error {
	return a.fs.Mkdir(a.Context(), name, perm)
}

// MkdirAll implements afero.Fs.
func (a *AccessorFs) MkdirAll(path string, perm fs.FileMode) error {
	return a.fs.MkdirAll(a.Context(), path, perm)
}

// Name implements afero.Fs.
func (a *AccessorFs) Name() string {
	return fmt.Sprintf("Context: %s", a.fs.Name())
}

// Open implements afero.Fs.
func (a *AccessorFs) Open(name string) (afero.File, error) {
	return a.fs.Open(a.Context(), name)
}

// OpenFile implements afero.Fs.
func (a *AccessorFs) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	return a.fs.OpenFile(a.Context(), name, flag, perm)
}

// Remove implements afero.Fs.
func (a *AccessorFs) Remove(name string) error {
	return a.fs.Remove(a.Context(), name)
}

// RemoveAll implements afero.Fs.
func (a *AccessorFs) RemoveAll(path string) error {
	return a.fs.RemoveAll(a.Context(), path)
}

// Rename implements afero.Fs.
func (a *AccessorFs) Rename(oldname string, newname string) error {
	return a.fs.Rename(a.Context(), oldname, newname)
}

// Stat implements afero.Fs.
func (a *AccessorFs) Stat(name string) (fs.FileInfo, error) {
	return a.fs.Stat(a.Context(), name)
}

var _ afero.Fs = (*AccessorFs)(nil)
