package user

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
	user   *github.User
}

// IsDir implements fs.FileInfo.
func (f *FileInfo) IsDir() bool {
	return true
}

// ModTime implements fs.FileInfo.
func (f *FileInfo) ModTime() time.Time {
	return f.user.GetUpdatedAt().Time
}

// Mode implements fs.FileInfo.
func (f *FileInfo) Mode() fs.FileMode {
	return os.ModeDir
}

// Name implements fs.FileInfo.
func (f *FileInfo) Name() string {
	return f.user.GetName()
}

// Size implements fs.FileInfo.
func (f *FileInfo) Size() int64 {
	panic("unimplemented")
}

// Sys implements fs.FileInfo.
func (f *FileInfo) Sys() any {
	return f.user
}

func Stat(ctx context.Context, gh *github.Client, name string) (*FileInfo, error) {
	user, _, err := gh.Users.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("stat user: %w", err)
	}

	return &FileInfo{
		client: gh,
		user:   user,
	}, nil
}
