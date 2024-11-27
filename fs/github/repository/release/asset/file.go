package asset

import (
	"context"
	"io"
	"io/fs"
	"net/http"
	"syscall"

	"github.com/google/go-github/v67/github"
	"github.com/unmango/go/fs/github/internal"
)

type File struct {
	internal.ReadOnlyFile
	internal.ReleasePath

	client *github.Client
	asset  *github.ReleaseAsset

	reader io.Reader
}

// Close implements afero.File.
func (f *File) Close() error {
	return nil
}

// Name implements afero.File.
func (f *File) Name() string {
	return f.asset.GetName()
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
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (f *File) Readdir(count int) ([]fs.FileInfo, error) {
	// TODO: Traverse into archives?
	panic("unimplemented")
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	// TODO: Traverse into archives?
	panic("unimplemented")
}

// Seek implements afero.File.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (f *File) Stat() (fs.FileInfo, error) {
	return &FileInfo{asset: f.asset}, nil
}

func (f *File) ensure() error {
	if f.reader != nil {
		return nil
	}

	reader, _, err := f.client.Repositories.DownloadReleaseAsset(
		context.TODO(),
		f.Owner,
		f.Repository,
		f.asset.GetID(),
		http.DefaultClient,
	)
	if err != nil {
		return err
	}

	f.reader = reader
	return nil
}

func NewFile(
	gh *github.Client,
	path internal.ReleasePath,
	asset *github.ReleaseAsset,
) *File {
	return &File{
		client:      gh,
		asset:       asset,
		ReleasePath: path,
	}
}
