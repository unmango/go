package docker

import (
	"context"
	"io"
	"io/fs"
	"syscall"

	"github.com/docker/docker/api/types/container"
	"github.com/moby/moby/client"
)

type File struct {
	client    client.ContainerAPIClient
	container string
	name      string

	stat   container.PathStat
	reader io.ReadCloser
}

// Close implements afero.File.
func (f *File) Close() error {
	if c, ok := f.reader.(io.ReadCloser); ok {
		return c.Close()
	}

	return nil
}

// Name implements afero.File.
func (f *File) Name() string {
	return f.name
}

// Read implements afero.File.
func (f *File) Read(p []byte) (n int, err error) {
	if err = f.ensure(); err != nil {
		return
	}

	return f.reader.Read(p)
}

// ReadAt implements afero.File.
func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	if ra, ok := f.reader.(io.ReaderAt); ok {
		return ra.ReadAt(p, off)
	}

	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (f *File) Readdir(count int) ([]fs.FileInfo, error) {
	panic("unimplemented")
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	panic("unimplemented")
}

// Seek implements afero.File.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	panic("unimplemented")
}

// Stat implements afero.File.
func (f *File) Stat() (fs.FileInfo, error) {
	return Stat(context.TODO(), f.client, f.container, f.name)
}

// Sync implements afero.File.
func (f *File) Sync() error {
	panic("unimplemented")
}

// Truncate implements afero.File.
func (f *File) Truncate(size int64) error {
	panic("unimplemented")
}

// Write implements afero.File.
func (f *File) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}

// WriteAt implements afero.File.
func (f *File) WriteAt(p []byte, off int64) (n int, err error) {
	panic("unimplemented")
}

// WriteString implements afero.File.
func (f *File) WriteString(s string) (ret int, err error) {
	panic("unimplemented")
}

func (f *File) ensure() error {
	if f.reader != nil {
		return nil
	}

	ctx := context.TODO()
	reader, stat, err := f.client.CopyFromContainer(ctx,
		f.container,
		f.name,
	)
	if err != nil {
		return err
	}

	f.reader = reader
	f.stat = stat

	return nil
}
