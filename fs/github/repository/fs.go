package repository

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
	"github.com/unmango/go/fs/github/internal"
	"github.com/unmango/go/fs/github/repository/release"
)

type Fs struct {
	afero.ReadOnlyFs
	internal.OwnerPath
	client *github.Client
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprint(f.OwnerPath)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	return Open(context.TODO(), f.client, f.OwnerPath, name)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return Open(context.TODO(), f.client, f.OwnerPath, name)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	return Stat(context.TODO(), f.client, f.OwnerPath, name)
}

func NewFs(gh *github.Client, owner string) afero.Fs {
	return &Fs{
		client:    gh,
		OwnerPath: internal.NewOwnerPath(owner),
	}
}

func Open(ctx context.Context, gh *github.Client, path internal.OwnerPath, name string) (afero.File, error) {
	p, err := path.Parse(name)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	repo, err := p.Repository()
	if err != nil {
		return nil, fmt.Errorf("invalid path %s: %w", path, err)
	}

	if _, err := p.Release(); err == nil {
		return release.Open(ctx, gh, internal.RepositoryPath{
			OwnerPath:  path,
			Repository: repo,
		}, name)
	}

	r, _, err := gh.Repositories.Get(ctx, path.Owner, name)
	if err != nil {
		return nil, err
	}

	return &File{
		client:    gh,
		repo:      r,
		OwnerPath: path,
	}, nil
}

func Readdir(ctx context.Context, gh *github.Client, user string, count int) ([]fs.FileInfo, error) {
	// TODO: Paging
	repos, _, err := gh.Repositories.ListByUser(ctx, user, nil)
	if err != nil {
		return nil, fmt.Errorf("user %s readdir: %w", user, err)
	}

	length := min(count, len(repos))
	results := make([]fs.FileInfo, length)

	for i := 0; i < length; i++ {
		results[i] = &FileInfo{repo: repos[i]}
	}

	return results, nil
}

func Readdirnames(ctx context.Context, gh *github.Client, user string, n int) ([]string, error) {
	// TODO: Paging
	repos, _, err := gh.Repositories.ListByUser(ctx, user, nil)
	if err != nil {
		return nil, fmt.Errorf("user %s readdirnames: %w", user, err)
	}

	length := min(n, len(repos))
	results := make([]string, length)

	for i := 0; i < length; i++ {
		results[i] = repos[i].GetName()
	}

	return results, nil
}

func Stat(ctx context.Context, gh *github.Client, path internal.OwnerPath, name string) (fs.FileInfo, error) {
	p, err := path.Parse(name)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	repo, err := p.Repository()
	if err != nil {
		return nil, fmt.Errorf("invalid path %s: %w", path, err)
	}

	if _, err := p.Release(); err == nil {
		return release.Stat(ctx, gh, internal.RepositoryPath{
			OwnerPath:  path,
			Repository: repo,
		}, name)
	}

	r, _, err := gh.Repositories.Get(ctx, path.Owner, name)
	if err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	return &FileInfo{repo: r}, nil
}
