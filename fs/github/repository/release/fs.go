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
	owner  string
	repo   string
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("https://github.com/%s/%s/releases", f.owner, f.repo)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	ctx := context.TODO()
	id, err := releaseId(ctx, f.client, f.owner, f.repo, name)
	if err != nil {
		return nil, err
	}

	return Open(ctx, f.client, f.owner, f.repo, id)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	ctx := context.TODO()
	id, err := releaseId(ctx, f.client, f.owner, f.repo, name)
	if err != nil {
		return nil, err
	}

	return Open(ctx, f.client, f.owner, f.repo, id)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	ctx := context.TODO()
	id, err := releaseId(ctx, f.client, f.owner, f.repo, name)
	if err != nil {
		return nil, err
	}

	return Stat(ctx, f.client, f.owner, f.repo, id)
}

func NewFs(gh *github.Client, owner, repository string) afero.Fs {
	return &Fs{
		client: gh,
		owner:  owner,
		repo:   repository,
	}
}

func Open(ctx context.Context, gh *github.Client, owner, repository string, id int64) (*File, error) {
	release, _, err := gh.Repositories.GetRelease(ctx, owner, repository, id)
	if err != nil {
		return nil, fmt.Errorf("open %d: %w", id, err)
	}

	return &File{
		client:  gh,
		release: release,
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

func Stat(ctx context.Context, gh *github.Client, owner, repository string, id int64) (*FileInfo, error) {
	release, _, err := gh.Repositories.GetRelease(ctx, owner, repository, id)
	if err != nil {
		return nil, fmt.Errorf("open %d: %w", id, err)
	}

	return &FileInfo{release: release}, nil
}

func releaseId(ctx context.Context, gh *github.Client, owner, repo, name string) (int64, error) {
	if id, err := strconv.ParseInt(name, 10, 64); err == nil {
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

	return 0, os.ErrNotExist
}
