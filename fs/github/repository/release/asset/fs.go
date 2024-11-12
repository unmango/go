package asset

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
	client  *github.Client
	owner   string
	repo    string
	release string
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("https://github.com/%s/%s/releases/%s/assets", f.owner, f.repo, f.release)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	id, err := f.assetId(name)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %w", name, err)
	}

	return Open(context.TODO(), f.client, f.owner, f.repo, id)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	id, err := f.assetId(name)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %w", name, err)
	}

	return Open(context.TODO(), f.client, f.owner, f.repo, id)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	id, err := f.assetId(name)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", name, err)
	}

	return Stat(context.TODO(), f.client, f.owner, f.repo, id)
}

func (f *Fs) assetId(name string) (int64, error) {
	ctx := context.TODO()
	releaseId, err := releaseId(ctx, f.client, f.owner, f.repo, f.release)
	if err != nil {
		return 0, fmt.Errorf("release %s: %w", f.release, err)
	}

	return assetId(ctx, f.client, f.owner, f.repo, releaseId, name)
}

func NewFs(gh *github.Client, owner, repository, release string) afero.Fs {
	return &Fs{
		client:  gh,
		release: release,
		owner:   owner,
		repo:    repository,
	}
}

func Open(ctx context.Context, gh *github.Client, owner, repository string, id int64) (*File, error) {
	asset, _, err := gh.Repositories.GetReleaseAsset(ctx, owner, repository, id)
	if err != nil {
		return nil, err
	}

	return &File{
		client: gh,
		owner:  owner,
		repo:   repository,
		asset:  asset,
	}, nil
}

func Readdir(ctx context.Context, gh *github.Client, owner, repository string, id int64, count int) ([]fs.FileInfo, error) {
	assets, _, err := gh.Repositories.ListReleaseAssets(ctx, owner, repository, id, nil)
	if err != nil {
		return nil, fmt.Errorf("%s/%s readdir: %w", owner, repository, err)
	}

	length := min(count, len(assets))
	results := make([]fs.FileInfo, length)

	for i := 0; i < length; i++ {
		results[i] = &FileInfo{asset: assets[i]}
	}

	return results, nil
}

func Readdirnames(ctx context.Context, gh *github.Client, owner, repository string, id int64, n int) ([]string, error) {
	assets, _, err := gh.Repositories.ListReleaseAssets(ctx, owner, repository, id, nil)
	if err != nil {
		return nil, fmt.Errorf("%s/%s readdir: %w", owner, repository, err)
	}

	length := min(n, len(assets))
	results := make([]string, length)

	for i := 0; i < length; i++ {
		results[i] = assets[i].GetName()
	}

	return results, nil
}

func Stat(ctx context.Context, gh *github.Client, owner, repository string, id int64) (*FileInfo, error) {
	asset, _, err := gh.Repositories.GetReleaseAsset(ctx, owner, repository, id)
	if err != nil {
		return nil, err
	}

	return &FileInfo{asset: asset}, nil
}

func releaseId(ctx context.Context, gh *github.Client, owner, repo, name string) (id int64, err error) {
	id, err = strconv.ParseInt(name, 10, 64)
	if err == nil {
		return
	}

	releases, _, err := gh.Repositories.ListReleases(ctx, owner, repo, nil)
	if err != nil {
		return 0, err
	}

	for _, r := range releases {
		if r.GetName() == name {
			return r.GetID(), nil
		}
	}

	return 0, os.ErrNotExist
}

func assetId(ctx context.Context, gh *github.Client, owner, repo string, releaseId int64, name string) (id int64, err error) {
	id, err = strconv.ParseInt(name, 10, 64)
	if err == nil {
		return
	}

	assets, _, err := gh.Repositories.ListReleaseAssets(ctx, owner, repo, releaseId, nil)
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
