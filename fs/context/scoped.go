package context

import (
	"fmt"
	"io/fs"
	"time"

	"github.com/spf13/afero"
	aferox "github.com/unmango/go/fs"
)

// ScopedFs adapts an [afero.Fs] to an [Fs] by calling [SetContext] on the given
// [ContextSetter] before forwarding each operation to the given [afero.Fs]
type ScopedFs struct {
	aferox.ContextSetter
	fs afero.Fs
}

// Chmod implements Fs.
func (s *ScopedFs) Chmod(ctx Context, name string, mode fs.FileMode) error {
	s.SetContext(ctx)
	return s.fs.Chmod(name, mode)
}

// Chown implements Fs.
func (s *ScopedFs) Chown(ctx Context, name string, uid int, gid int) error {
	s.SetContext(ctx)
	return s.fs.Chown(name, uid, gid)
}

// Chtimes implements Fs.
func (s *ScopedFs) Chtimes(ctx Context, name string, atime time.Time, mtime time.Time) error {
	s.SetContext(ctx)
	return s.fs.Chtimes(name, atime, mtime)
}

// Create implements Fs.
func (s *ScopedFs) Create(ctx Context, name string) (afero.File, error) {
	s.SetContext(ctx)
	return s.fs.Create(name)
}

// Mkdir implements Fs.
func (s *ScopedFs) Mkdir(ctx Context, name string, perm fs.FileMode) error {
	s.SetContext(ctx)
	return s.fs.Mkdir(name, perm)
}

// MkdirAll implements Fs.
func (s *ScopedFs) MkdirAll(ctx Context, path string, perm fs.FileMode) error {
	s.SetContext(ctx)
	return s.fs.MkdirAll(path, perm)
}

// Name implements Fs.
func (s *ScopedFs) Name() string {
	return fmt.Sprintf("Scoped: %s", s.fs.Name())
}

// Open implements Fs.
func (s *ScopedFs) Open(ctx Context, name string) (afero.File, error) {
	s.SetContext(ctx)
	return s.fs.Open(name)
}

// OpenFile implements Fs.
func (s *ScopedFs) OpenFile(ctx Context, name string, flag int, perm fs.FileMode) (afero.File, error) {
	s.SetContext(ctx)
	return s.fs.OpenFile(name, flag, perm)
}

// Remove implements Fs.
func (s *ScopedFs) Remove(ctx Context, name string) error {
	s.SetContext(ctx)
	return s.fs.Remove(name)
}

// RemoveAll implements Fs.
func (s *ScopedFs) RemoveAll(ctx Context, path string) error {
	s.SetContext(ctx)
	return s.fs.RemoveAll(path)
}

// Rename implements Fs.
func (s *ScopedFs) Rename(ctx Context, oldname string, newname string) error {
	s.SetContext(ctx)
	return s.fs.Rename(oldname, newname)
}

// Stat implements Fs.
func (s *ScopedFs) Stat(ctx Context, name string) (fs.FileInfo, error) {
	s.SetContext(ctx)
	return s.fs.Stat(name)
}

var _ Fs = (*ScopedFs)(nil)
