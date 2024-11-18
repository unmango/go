package docker

import (
	"context"
	"fmt"
	"io/fs"
	"time"

	"github.com/docker/docker/client"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/docker/internal"
)

type Fs struct {
	client    client.ContainerAPIClient
	container string
}

// Chmod implements afero.Fs.
func (f Fs) Chmod(name string, mode fs.FileMode) error {
	return f.exec(context.TODO(), "chmod", mode.String(), name)
}

// Chown implements afero.Fs.
func (f Fs) Chown(name string, uid int, gid int) error {
	return f.exec(context.TODO(),
		"chown", fmt.Sprintf("%d:%d", uid, gid), name,
	)
}

// Chtimes implements afero.Fs.
func (f Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	panic("unimplemented")
}

// Create implements afero.Fs.
func (f Fs) Create(name string) (afero.File, error) {
	err := f.exec(context.TODO(), "touch", name)
	if err != nil {
		return nil, err
	}

	// TODO: Less lazy?
	return &File{
		client:    f.client,
		container: f.container,
		name:      name,
	}, nil
}

// Mkdir implements afero.Fs.
func (f Fs) Mkdir(name string, perm fs.FileMode) error {
	return f.exec(context.TODO(),
		"mkdir", fmt.Sprintf("--mode=%d", perm), name,
	)
}

// MkdirAll implements afero.Fs.
func (f Fs) MkdirAll(path string, perm fs.FileMode) error {
	return f.exec(context.TODO(),
		"mkdir", "--parents", fmt.Sprintf("--mode=%d", perm), path,
	)
}

// Name implements afero.Fs.
func (f Fs) Name() string {
	return f.container
}

// Open implements afero.Fs.
func (f Fs) Open(name string) (afero.File, error) {
	// TODO: Less lazy?
	return &File{
		client:    f.client,
		container: f.container,
		name:      name,
	}, nil
}

// OpenFile implements afero.Fs.
func (f Fs) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	// TODO: Actual implementation
	// TODO: Less lazy?
	return &File{
		client:    f.client,
		container: f.container,
		name:      name,
	}, nil
}

// Remove implements afero.Fs.
func (f Fs) Remove(name string) error {
	return f.exec(context.TODO(), "rm", name)
}

// RemoveAll implements afero.Fs.
func (f Fs) RemoveAll(path string) error {
	return f.exec(context.TODO(), "rm", "--recursive", path)
}

// Rename implements afero.Fs.
func (f Fs) Rename(oldname string, newname string) error {
	return f.exec(context.TODO(), "mv", oldname, newname)
}

// Stat implements afero.Fs.
func (f Fs) Stat(name string) (fs.FileInfo, error) {
	return Stat(context.TODO(), f.client, f.container, name)
}

func (f Fs) exec(ctx context.Context, cmd ...string) error {
	return internal.Exec(ctx, f.client, f.container,
		internal.ExecOptions{
			Cmd: cmd,
		},
	)
}

func NewFs(client client.ContainerAPIClient, container string) afero.Fs {
	return Fs{client, container}
}
