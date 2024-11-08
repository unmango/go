package repository

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
)

type repository struct {
	*github.RepositoriesService

	owner, name string
}

// Chmod implements afero.Fs.
func (r *repository) Chmod(name string, mode fs.FileMode) error { panic("unsupported") }

// Chown implements afero.Fs.
func (r *repository) Chown(name string, uid int, gid int) error { panic("unsupported") }

// Chtimes implements afero.Fs.
func (r *repository) Chtimes(name string, atime time.Time, mtime time.Time) error {
	panic("unsupported")
}

// Create implements afero.Fs.
func (r *repository) Create(name string) (afero.File, error) {
	return &fileWriter{name, r.RepositoriesService}, nil
}

// Mkdir implements afero.Fs.
func (r *repository) Mkdir(name string, perm fs.FileMode) error {
	return errors.New("unsupported operation")
}

// MkdirAll implements afero.Fs.
func (r *repository) MkdirAll(path string, perm fs.FileMode) error {
	return errors.New("unsupported operation")
}

// Name implements afero.Fs.
func (r *repository) Name() string { return fmt.Sprintf("%s/%s", r.owner, r.name) }

// Open implements afero.Fs.
func (r *repository) Open(name string) (afero.File, error) {
	ctx := context.Background()
	fc, _, _, err := r.GetContents(ctx, r.owner, r.name, name, nil)
	if err != nil {
		return nil, err
	}
	if fc == nil {
		return nil, fmt.Errorf("not a file: %s", name)
	}

	return &fileContent{file: fc}, nil
}

// OpenFile implements afero.Fs.
func (r *repository) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	ctx := context.Background()
	fc, _, _, err := r.GetContents(ctx, r.owner, r.name, name, nil)
	if err != nil {
		return nil, err
	}
	if fc == nil {
		return nil, fmt.Errorf("not a file: %s", name)
	}

	return &fileContent{file: fc}, nil
}

// Remove implements afero.Fs.
func (r *repository) Remove(name string) error {
	panic("unimplemented")
}

// RemoveAll implements afero.Fs.
func (r *repository) RemoveAll(path string) error {
	panic("unimplemented")
}

// Rename implements afero.Fs.
func (r *repository) Rename(oldname string, newname string) error {
	panic("unimplemented")
}

// Stat implements afero.Fs.
func (r *repository) Stat(name string) (fs.FileInfo, error) {
	ctx := context.Background()
	fc, dc, _, err := r.GetContents(ctx, r.owner, r.name, name, nil)
	if err != nil {
		return nil, err
	}

	if fc != nil {
		return &fileInfoContent{file: fc}, nil
	}
	if dc != nil {
		return &directoryContent{name, dc}, nil
	}

	return nil, fmt.Errorf("the underlying GitHub client misbehaved: %s",
		"both file content and directory content were nil",
	)
}

var _ afero.Fs = &repository{}

func New(owner, name string, client *github.RepositoriesService) *repository {
	return &repository{client, owner, name}
}
