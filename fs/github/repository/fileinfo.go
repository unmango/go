package repository

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/google/go-github/v66/github"
)

type FileInfo struct {
	client *github.Client
	repo   *github.Repository
}

// IsDir implements fs.FileInfo.
func (f *FileInfo) IsDir() bool {
	return true
}

// ModTime implements fs.FileInfo.
func (f *FileInfo) ModTime() time.Time {
	return f.repo.GetUpdatedAt().Time
}

// Mode implements fs.FileInfo.
func (f *FileInfo) Mode() fs.FileMode {
	return os.ModeDir
}

// Name implements fs.FileInfo.
func (f *FileInfo) Name() string {
	return f.repo.GetName()
}

// Size implements fs.FileInfo.
func (f *FileInfo) Size() int64 {
	return int64(f.repo.GetSize())
}

// Sys implements fs.FileInfo.
func (f *FileInfo) Sys() any {
	return f.repo
}

func Stat(ctx context.Context, gh *github.Client, user, name string) (*FileInfo, error) {
	repo, _, err := gh.Repositories.Get(ctx, user, name)
	if err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	return &FileInfo{
		client: gh,
		repo:   repo,
	}, nil
}
