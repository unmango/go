package release

import (
	"io"
	"io/fs"
	"net/http"
	"syscall"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
)

type assetFile struct {
	internal.Fs
	client *github.RepositoriesService
	asset  *github.ReleaseAsset

	reader io.ReadCloser
}

// Close implements afero.File.
func (a *assetFile) Close() error {
	if a.reader == nil {
		return nil
	}

	return a.reader.Close()
}

// Name implements afero.File.
func (a *assetFile) Name() string {
	return a.asset.GetName()
}

// Read implements afero.File.
func (a *assetFile) Read(p []byte) (n int, err error) {
	if err = a.ensureReader(); err != nil {
		return
	}

	return a.reader.Read(p)
}

// ReadAt implements afero.File.
func (a *assetFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (a *assetFile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, syscall.EPERM
}

// Readdirnames implements afero.File.
func (a *assetFile) Readdirnames(n int) ([]string, error) {
	return nil, syscall.EPERM
}

// Seek implements afero.File.
func (a *assetFile) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (a *assetFile) Stat() (fs.FileInfo, error) {
	return &assetFileInfo{
		asset: a.asset,
	}, nil
}

// Sync implements afero.File.
func (a *assetFile) Sync() error {
	return syscall.EPERM
}

// Truncate implements afero.File.
func (a *assetFile) Truncate(size int64) error {
	return syscall.EPERM
}

// Write implements afero.File.
func (a *assetFile) Write(p []byte) (n int, err error) {
	return 0, syscall.EPERM
}

// WriteAt implements afero.File.
func (a *assetFile) WriteAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// WriteString implements afero.File.
func (a *assetFile) WriteString(s string) (ret int, err error) {
	return 0, syscall.EPERM
}

func (a *assetFile) ensureReader() error {
	if a.reader != nil {
		return nil
	}

	reader, _, err := a.client.DownloadReleaseAsset(
		a.Context(),
		a.Owner,
		a.Repo,
		a.asset.GetID(),
		http.DefaultClient,
	)
	if err != nil {
		return err
	}

	a.reader = reader
	return nil
}
