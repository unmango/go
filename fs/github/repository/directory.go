package repository

import (
	"io/fs"
	"time"

	"github.com/google/go-github/v66/github"
)

type directoryContent struct {
	name  string
	files []*github.RepositoryContent
}

// IsDir implements fs.FileInfo.
func (*directoryContent) IsDir() bool { return true }

// ModTime implements fs.FileInfo.
func (*directoryContent) ModTime() time.Time { panic("unsupported") }

// Mode implements fs.FileInfo.
func (*directoryContent) Mode() fs.FileMode { panic("unsupported") }

// Name implements fs.FileInfo.
func (r *directoryContent) Name() string { return r.name }

// Size implements fs.FileInfo.
func (r *directoryContent) Size() (s int64) {
	for _, f := range r.files {
		s += int64(f.GetSize())
	}

	return
}

// Sys implements fs.FileInfo.
func (r *directoryContent) Sys() any { return r.files }

func NewDirectoryFileInfo(name string, contents []*github.RepositoryContent) fs.FileInfo {
	return &directoryContent{name, contents}
}
