package release

import (
	"context"
	"io/fs"
	"syscall"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
)

type File struct {
	internal.ReadOnlyFile
	asset *github.ReleaseAsset
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
	panic("unimplemented")
}

// ReadAt implements afero.File.
func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
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
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (f *File) Stat() (fs.FileInfo, error) {
	return &FileInfo{asset: f.asset}, nil
}

func Open(ctx context.Context, gh *github.Client, owner, repository string, id int64) (*File, error) {
	asset, _, err := gh.Repositories.GetReleaseAsset(ctx, owner, repository, id)
	if err != nil {
		return nil, err
	}

	return &File{asset: asset}, nil
}
