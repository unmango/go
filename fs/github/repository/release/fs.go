package release

import (
	"context"
	"fmt"
	"io/fs"
	"os"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/internal"
	"github.com/unmango/go/fs/github/repository/release/asset"
)

type Fs struct {
	afero.ReadOnlyFs
	internal.RepositoryPath
	client *github.Client
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("%s/releases", f.RepositoryPath)
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

func NewFs(gh *github.Client, owner, repository string) afero.Fs {
	return &Fs{
		client:         gh,
		RepositoryPath: internal.NewRepositoryPath(owner, repository),
	}
}

func Open(ctx context.Context, gh *github.Client, path internal.Path) (afero.File, error) {
	release, err := internal.ParseRelease(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path %s: %w", path, err)
	}

	if _, err := path.Asset(); err == nil {
		return asset.Open(ctx, gh, path)
	}

	id, err := releaseId(ctx, gh, release)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}

	r, _, err := gh.Repositories.GetRelease(ctx, release.Owner, release.Repository, id)
	if err != nil {
		return nil, fmt.Errorf("open %d: %w", id, err)
	}

	return &File{
		client:         gh,
		release:        r,
		RepositoryPath: release.RepositoryPath,
	}, nil
}

func Readdir(ctx context.Context, gh *github.Client, owner, repository string, count int) ([]fs.FileInfo, error) {
	releases, _, err := gh.Repositories.ListReleases(ctx, owner, repository, nil)
	if err != nil {
		return nil, fmt.Errorf("%s/%s readdir: %w", owner, repository, err)
	}

	length := min(count, len(releases))
	results := make([]fs.FileInfo, length)

	for i := 0; i < length; i++ {
		results[i] = &FileInfo{release: releases[i]}
	}

	return results, nil
}

func Readdirnames(ctx context.Context, gh *github.Client, owner, repository string, n int) ([]string, error) {
	releases, _, err := gh.Repositories.ListReleases(ctx, owner, repository, nil)
	if err != nil {
		return nil, fmt.Errorf("%s/%s readdir: %w", owner, repository, err)
	}

	length := min(n, len(releases))
	results := make([]string, length)

	for i := 0; i < length; i++ {
		results[i] = releases[i].GetName()
	}

	return results, nil
}

func Stat(ctx context.Context, gh *github.Client, path internal.Path) (fs.FileInfo, error) {
	release, err := internal.ParseRelease(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path %s: %w", release, err)
	}

	if _, err := path.Asset(); err == nil {
		return asset.Stat(ctx, gh, path)
	}

	id, err := releaseId(ctx, gh, release)
	if err != nil {
		return nil, fmt.Errorf("reading release id: %w", err)
	}

	r, _, err := gh.Repositories.GetRelease(ctx, release.Owner, release.Repository, id)
	if err != nil {
		return nil, fmt.Errorf("open %d: %w", id, err)
	}

	return &FileInfo{release: r}, nil
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

	return 0, os.ErrNotExist
}
