package asset

import (
	"context"
	"fmt"
	"io/fs"
	"os"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/internal"
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
	return Open(context.TODO(), f.client, f.owner, f.repo, f.release, name)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return Open(context.TODO(), f.client, f.owner, f.repo, f.release, name)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	return Stat(context.TODO(), f.client, f.owner, f.repo, f.release, name)
}

func NewFs(gh *github.Client, owner, repository, release string) afero.Fs {
	return &Fs{
		client:  gh,
		release: release,
		owner:   owner,
		repo:    repository,
	}
}

func Open(
	ctx context.Context,
	gh *github.Client,
	owner, repository string,
	release, name string,
) (*File, error) {
	id, err := assetId(ctx, gh, owner, repository, release, name)
	if err != nil {
		return nil, fmt.Errorf("opening %s/%s/%s/%s: %w", owner, repository, release, name, err)
	}

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

func Stat(ctx context.Context, gh *github.Client, owner, repository, release, name string) (*FileInfo, error) {
	id, err := assetId(ctx, gh, owner, repository, release, name)
	if err != nil {
		return nil, fmt.Errorf("reading asset id: %w", err)
	}

	asset, _, err := gh.Repositories.GetReleaseAsset(ctx, owner, repository, id)
	if err != nil {
		return nil, err
	}

	return &FileInfo{asset: asset}, nil
}

func releaseId(ctx context.Context, gh *github.Client, owner, repo, name string) (int64, error) {
	if id, ok := internal.TryGetId(name); ok {
		return id, nil
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

	return 0, fmt.Errorf("%s: %w", name, os.ErrNotExist)
}

func assetId(ctx context.Context, gh *github.Client, owner, repo, release string, name string) (int64, error) {
	releaseId, err := releaseId(ctx, gh, owner, repo, release)
	if err != nil {
		return 0, err
	}

	if id, ok := internal.TryGetId(name); ok {
		return id, nil
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

	return 0, fmt.Errorf("%s: %w", name, os.ErrNotExist)
}
