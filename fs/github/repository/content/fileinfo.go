package content

import (
	"io/fs"
	"time"

	"github.com/google/go-github/v66/github"
)

type FileInfo struct {
	content *github.RepositoryContent
}

// IsDir implements fs.FileInfo.
func (f *FileInfo) IsDir() bool {
	return false
}

// ModTime implements fs.FileInfo.
func (f *FileInfo) ModTime() time.Time {
	panic("unimplemented")
}

// Mode implements fs.FileInfo.
func (f *FileInfo) Mode() fs.FileMode {
	return 0
}

// Name implements fs.FileInfo.
func (f *FileInfo) Name() string {
	return f.content.GetName()
}

// Size implements fs.FileInfo.
func (f *FileInfo) Size() int64 {
	return int64(f.content.GetSize())
}

// Sys implements fs.FileInfo.
func (f *FileInfo) Sys() any {
	return f.content
}
