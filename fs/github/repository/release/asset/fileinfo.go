package asset

import (
	"io/fs"
	"time"

	"github.com/google/go-github/v67/github"
)

type FileInfo struct {
	asset *github.ReleaseAsset
}

// IsDir implements fs.FileInfo.
func (a *FileInfo) IsDir() bool {
	return false
}

// ModTime implements fs.FileInfo.
func (a *FileInfo) ModTime() time.Time {
	return a.asset.GetUpdatedAt().Time
}

// Mode implements fs.FileInfo.
func (a *FileInfo) Mode() fs.FileMode {
	return 0
}

// Name implements fs.FileInfo.
func (a *FileInfo) Name() string {
	return a.asset.GetName()
}

// Size implements fs.FileInfo.
func (a *FileInfo) Size() int64 {
	return int64(a.asset.GetSize())
}

// Sys implements fs.FileInfo.
func (a *FileInfo) Sys() any {
	return a.asset
}
