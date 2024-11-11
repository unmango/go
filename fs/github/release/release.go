package release

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/internal"
)

type release struct {
	internal.Fs
	name   string
	client *github.RepositoriesService
}

// Chmod implements afero.Fs.
func (*release) Chmod(name string, mode fs.FileMode) error {
	return syscall.EPERM
}

// Chown implements afero.Fs.
func (*release) Chown(name string, uid int, gid int) error {
	return syscall.EPERM
}

// Chtimes implements afero.Fs.
func (*release) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return syscall.EPERM
}

// Create implements afero.Fs.
func (*release) Create(name string) (afero.File, error) {
	return nil, syscall.EPERM
}

// Mkdir implements afero.Fs.
func (*release) Mkdir(name string, perm fs.FileMode) error {
	return syscall.EPERM
}

// MkdirAll implements afero.Fs.
func (*release) MkdirAll(path string, perm fs.FileMode) error {
	return syscall.EPERM
}

// Name implements afero.Fs.
func (r *release) Name() string {
	return r.name
}

// Open implements afero.Fs.
func (r *release) Open(name string) (afero.File, error) {
	asset, err := r.getAsset(name)
	if err != nil {
		return nil, err
	}

	return &assetFile{
		client: r.client,
		asset:  asset,
		Fs:     r.Fs,
	}, nil
}

// OpenFile implements afero.Fs.
func (r *release) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return r.Open(name)
}

// Remove implements afero.Fs.
func (r *release) Remove(name string) error {
	return syscall.EPERM
}

// RemoveAll implements afero.Fs.
func (r *release) RemoveAll(path string) error {
	return syscall.EPERM
}

// Rename implements afero.Fs.
func (r *release) Rename(oldname string, newname string) error {
	return syscall.EPERM
}

// Stat implements afero.Fs.
func (r *release) Stat(name string) (fs.FileInfo, error) {
	asset, err := r.getAsset(name)
	if err != nil {
		return nil, err
	}

	return &assetFileInfo{asset}, nil
}

func (r *release) getAsset(name string) (*github.ReleaseAsset, error) {
	id, err := strconv.ParseInt(filepath.Base(name), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid release asset id: %w", err)
	}

	asset, _, err := r.client.GetReleaseAsset(r.Context(), r.Owner, r.Repo, id)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (r *release) getAssetId(name string) (id int64, err error) {
	id, err = strconv.ParseInt(filepath.Base(name), 10, 64)
	if err == nil {
		return
	}

	ctx := r.Context()
	assets, _, err := r.client.ListReleaseAssets(
		ctx,
		r.Owner,
		r.Repo,
		69,
		nil,
	)
	if err != nil {
		return
	}

	for _, a := range assets {
		if a.GetName() == name {
			return a.GetID(), nil
		}
	}

	err = fmt.Errorf("asset not found: %s", name)
	return
}

func New(owner, repo, name string, client *github.RepositoriesService) afero.Fs {
	return &release{
		Fs: internal.Fs{
			ContextAccessor: internal.BackgroundContext(),
			Owner:           owner,
			Repo:            repo,
		},
		name:   name,
		client: client,
	}
}
