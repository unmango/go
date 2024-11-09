package release

import (
	"io/fs"
	"time"

	"github.com/google/go-github/v66/github"
)

type assetFileInfo struct {
	asset *github.ReleaseAsset
}

// IsDir implements fs.FileInfo.
func (a *assetFileInfo) IsDir() bool {
	return false
}

// ModTime implements fs.FileInfo.
func (a *assetFileInfo) ModTime() time.Time {
	panic("unsupported")
}

// Mode implements fs.FileInfo.
func (a *assetFileInfo) Mode() fs.FileMode {
	panic("unsupported")
}

// Name implements fs.FileInfo.
func (a *assetFileInfo) Name() string {
	return a.asset.GetName()
}

// Size implements fs.FileInfo.
func (a *assetFileInfo) Size() int64 {
	return int64(a.asset.GetSize())
}

// Sys implements fs.FileInfo.
func (a *assetFileInfo) Sys() any {
	return a.asset
}
