package context

import (
	"context"
	"io/fs"
	"time"

	"github.com/spf13/afero"
)

type Adapter struct{ Fs }

// ChmodContext implements AferoFs.
func (a *Adapter) ChmodContext(ctx context.Context, name string, mode fs.FileMode) error {
	return a.Chmod(ctx, name, mode)
}

// ChownContext implements AferoFs.
func (a *Adapter) ChownContext(ctx context.Context, name string, uid int, gid int) error {
	panic("unimplemented")
}

// ChtimesContext implements AferoFs.
func (a *Adapter) ChtimesContext(ctx context.Context, name string, atime time.Time, mtime time.Time) error {
	panic("unimplemented")
}

// CreateContext implements AferoFs.
func (a *Adapter) CreateContext(ctx context.Context, name string) (afero.File, error) {
	panic("unimplemented")
}

// MkdirAllContext implements AferoFs.
func (a *Adapter) MkdirAllContext(ctx context.Context, path string, perm fs.FileMode) error {
	panic("unimplemented")
}

// MkdirContext implements AferoFs.
func (a *Adapter) MkdirContext(ctx context.Context, name string, perm fs.FileMode) error {
	panic("unimplemented")
}

// OpenContext implements AferoFs.
func (a *Adapter) OpenContext(ctx context.Context, name string) (afero.File, error) {
	panic("unimplemented")
}

// OpenFileContext implements AferoFs.
func (a *Adapter) OpenFileContext(ctx context.Context, name string, flag int, perm fs.FileMode) (afero.File, error) {
	panic("unimplemented")
}

// RemoveAllContext implements AferoFs.
func (a *Adapter) RemoveAllContext(ctx context.Context, path string) error {
	panic("unimplemented")
}

// RemoveContext implements AferoFs.
func (a *Adapter) RemoveContext(ctx context.Context, name string) error {
	panic("unimplemented")
}

// RenameContext implements AferoFs.
func (a *Adapter) RenameContext(ctx context.Context, oldname string, newname string) error {
	panic("unimplemented")
}

// StatContext implements AferoFs.
func (a *Adapter) StatContext(ctx context.Context, name string) (fs.FileInfo, error) {
	panic("unimplemented")
}

type DiscardAferoFs struct{ afero.Fs }

// ChmodContext implements AferoFs.
func (a *DiscardAferoFs) ChmodContext(ctx context.Context, name string, mode fs.FileMode) error {
	return a.Chmod(name, mode)
}

// ChownContext implements AferoFs.
func (a *DiscardAferoFs) ChownContext(ctx context.Context, name string, uid int, gid int) error {
	return a.Chown(name, uid, gid)
}

// ChtimesContext implements AferoFs.
func (a *DiscardAferoFs) ChtimesContext(ctx context.Context, name string, atime time.Time, mtime time.Time) error {
	return a.Chtimes(name, atime, mtime)
}

// CreateContext implements AferoFs.
func (a *DiscardAferoFs) CreateContext(ctx context.Context, name string) (afero.File, error) {
	return a.Create(name)
}

// MkdirAllContext implements AferoFs.
func (a *DiscardAferoFs) MkdirAllContext(ctx context.Context, path string, perm fs.FileMode) error {
	return a.MkdirAll(path, perm)
}

// MkdirContext implements AferoFs.
func (a *DiscardAferoFs) MkdirContext(ctx context.Context, name string, perm fs.FileMode) error {
	return a.Mkdir(name, perm)
}

// OpenContext implements AferoFs.
func (a *DiscardAferoFs) OpenContext(ctx context.Context, name string) (afero.File, error) {
	return a.Open(name)
}

// OpenFileContext implements AferoFs.
func (a *DiscardAferoFs) OpenFileContext(ctx context.Context, name string, flag int, perm fs.FileMode) (afero.File, error) {
	return a.OpenFile(name, flag, perm)
}

// RemoveAllContext implements AferoFs.
func (a *DiscardAferoFs) RemoveAllContext(ctx context.Context, path string) error {
	return a.RemoveAll(path)
}

// RemoveContext implements AferoFs.
func (a *DiscardAferoFs) RemoveContext(ctx context.Context, name string) error {
	return a.Remove(name)
}

// RenameContext implements AferoFs.
func (a *DiscardAferoFs) RenameContext(ctx context.Context, oldname string, newname string) error {
	return a.Rename(oldname, newname)
}

// StatContext implements AferoFs.
func (a *DiscardAferoFs) StatContext(ctx context.Context, name string) (fs.FileInfo, error) {
	return a.Stat(name)
}

var _ AferoFs = (*DiscardAferoFs)(nil)
