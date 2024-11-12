package release

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"strconv"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
)

type Fs struct {
	afero.ReadOnlyFs
	client *github.Client
	name   string
	owner  string
	repo   string
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("https://github.com/%s/%s/releases/%s", f.owner, f.repo, f.name)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	asset, err := f.asset(name)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}

	return &File{asset: asset}, nil
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	asset, err := f.asset(name)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %w", name, err)
	}

	return &File{asset: asset}, nil
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	asset, err := f.asset(name)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", name, err)
	}

	return &FileInfo{asset: asset}, nil
}

func (f *Fs) asset(name string) (*github.ReleaseAsset, error) {
	releaseId, err := f.releaseId()
	if err != nil {
		return nil, fmt.Errorf("release %s: %w", f.name, err)
	}

	assetId, err := f.assetId(releaseId, name)
	if err != nil {
		return nil, err
	}

	asset, _, err := f.client.Repositories.GetReleaseAsset(
		context.TODO(),
		f.owner,
		f.name,
		assetId,
	)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (f *Fs) assetId(releaseId int64, name string) (id int64, err error) {
	id, err = strconv.ParseInt(f.name, 10, 64)
	if err == nil {
		return
	}

	assets, _, err := f.client.Repositories.ListReleaseAssets(
		context.TODO(),
		f.owner,
		f.repo,
		releaseId,
		nil,
	)
	if err != nil {
		return 0, err
	}

	for _, a := range assets {
		if a.GetName() == name {
			return a.GetID(), nil
		}
	}

	return 0, os.ErrNotExist
}

func (f *Fs) releaseId() (id int64, err error) {
	id, err = strconv.ParseInt(f.name, 10, 64)
	if err == nil {
		return
	}

	releases, _, err := f.client.Repositories.ListReleases(
		context.TODO(),
		f.owner,
		f.repo,
		nil,
	)
	if err != nil {
		return 0, err
	}

	for _, r := range releases {
		if r.GetName() == f.name {
			return r.GetID(), nil
		}
	}

	return 0, os.ErrNotExist
}

func New(gh *github.Client, owner, repository, name string) afero.Fs {
	return &Fs{
		client: gh,
		name:   name,
		owner:  owner,
		repo:   repository,
	}
}
