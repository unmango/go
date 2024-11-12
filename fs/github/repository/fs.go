package repository

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/google/go-github/v66/github"
	"github.com/spf13/afero"
)

type Fs struct {
	afero.ReadOnlyFs
	client *github.Client
	owner  string
}

// Name implements afero.Fs.
func (f *Fs) Name() string {
	return fmt.Sprintf("https://github.com/%s", f.owner)
}

// Open implements afero.Fs.
func (f *Fs) Open(name string) (afero.File, error) {
	return Open(context.TODO(), f.client, f.owner, name)
}

// OpenFile implements afero.Fs.
func (f *Fs) OpenFile(name string, _ int, _ fs.FileMode) (afero.File, error) {
	return Open(context.TODO(), f.client, f.owner, name)
}

// Stat implements afero.Fs.
func (f *Fs) Stat(name string) (fs.FileInfo, error) {
	return Stat(context.TODO(), f.client, f.owner, name)
}

func New(gh *github.Client, owner string) afero.Fs {
	return &Fs{client: gh, owner: owner}
}

func Open(ctx context.Context, gh *github.Client, user, name string) (*File, error) {
	repo, _, err := gh.Repositories.Get(ctx, user, name)
	if err != nil {
		return nil, err
	}

	return &File{
		client: gh,
		repo:   repo,
	}, nil
}

func Readdir(ctx context.Context, gh *github.Client, user string, count int) ([]fs.FileInfo, error) {
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

func Stat(ctx context.Context, gh *github.Client, user, name string) (*FileInfo, error) {
	repo, _, err := gh.Repositories.Get(ctx, user, name)
	if err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	return &FileInfo{repo: repo}, nil
}
