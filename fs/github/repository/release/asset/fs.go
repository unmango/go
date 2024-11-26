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
	internal.ReleasePath
	client *github.Client
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("%s/download", f.ReleasePath)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	path, err := f.Parse(name)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}

	return Open(context.TODO(), f.client, path)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	path, err := f.Parse(name)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}

	return Open(context.TODO(), f.client, path)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	path, err := f.Parse(name)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", name, err)
	}

	return Stat(context.TODO(), f.client, path)
}

func NewFs(gh *github.Client, owner, repository, release string) afero.Fs {
	return &Fs{
		client:      gh,
		ReleasePath: internal.NewReleasePath(owner, repository, release),
	}
}

func Open(ctx context.Context, gh *github.Client, path internal.Path) (*File, error) {
	assetPath, err := internal.ParseAsset(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path %s: %w", path, err)
	}

	id, err := assetId(ctx, gh, assetPath)
	if err != nil {
		return nil, fmt.Errorf("opening %s: %w", path, err)
	}

	asset, _, err := gh.Repositories.GetReleaseAsset(ctx, assetPath.Owner, assetPath.Repository, id)
	if err != nil {
		return nil, err
	}

	return &File{
		client:      gh,
		asset:       asset,
		ReleasePath: assetPath.ReleasePath,
	}, nil
}

func Readdir(ctx context.Context, gh *github.Client, path internal.RepositoryPath, id int64, count int) ([]fs.FileInfo, error) {
	assets, _, err := gh.Repositories.ListReleaseAssets(ctx, path.Owner, path.Repository, id, nil)
	if err != nil {
		return nil, fmt.Errorf("readdir %s: %w", path, err)
	}

	length := min(count, len(assets))
	results := make([]fs.FileInfo, length)

	for i := 0; i < length; i++ {
		results[i] = &FileInfo{asset: assets[i]}
	}

	return results, nil
}

func Readdirnames(ctx context.Context, gh *github.Client, path internal.RepositoryPath, id int64, n int) ([]string, error) {
	assets, _, err := gh.Repositories.ListReleaseAssets(ctx, path.Owner, path.Repository, id, nil)
	if err != nil {
		return nil, fmt.Errorf("readdir %s: %w", path, err)
	}

	length := min(n, len(assets))
	results := make([]string, length)

	for i := 0; i < length; i++ {
		results[i] = assets[i].GetName()
	}

	return results, nil
}

func Stat(ctx context.Context, gh *github.Client, path internal.Path) (*FileInfo, error) {
	assetPath, err := internal.ParseAsset(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path %s: %w", path, err)
	}

	id, err := assetId(ctx, gh, assetPath)
	if err != nil {
		return nil, fmt.Errorf("reading asset id: %w", err)
	}

	asset, _, err := gh.Repositories.GetReleaseAsset(ctx, assetPath.Owner, assetPath.Repository, id)
	if err != nil {
		return nil, err
	}

	return &FileInfo{asset: asset}, nil
}

func releaseId(ctx context.Context, gh *github.Client, path internal.ReleasePath) (int64, error) {
	if id, ok := internal.TryGetId(path.Release); ok {
		return id, nil
	}

	releases, _, err := gh.Repositories.ListReleases(ctx, path.Owner, path.Repository, nil)
	if err != nil {
		return 0, err
	}

	for _, r := range releases {
		if r.GetName() == path.Release {
			return r.GetID(), nil
		}
	}

	return 0, fmt.Errorf("%s: %w", path.Release, os.ErrNotExist)
}

func assetId(ctx context.Context, gh *github.Client, path internal.AssetPath) (int64, error) {
	if id, ok := internal.TryGetId(path.Asset); ok {
		return id, nil
	}

	releaseId, err := releaseId(ctx, gh, path.ReleasePath)
	if err != nil {
		return 0, fmt.Errorf("reading release id: %w", err)
	}

	assets, _, err := gh.Repositories.ListReleaseAssets(ctx, path.Owner, path.Repository, releaseId, nil)
	if err != nil {
		return 0, err
	}

	for _, a := range assets {
		if a.GetName() == path.Asset {
			return a.GetID(), nil
		}
	}

	return 0, fmt.Errorf("%s: %w", path, os.ErrNotExist)
}
