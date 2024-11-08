package repository

import (
	"io/fs"
	"time"

	"github.com/google/go-github/v66/github"
)

type fileInfoContent struct {
	file *github.RepositoryContent
}

// IsDir implements fs.FileInfo.
func (r *fileInfoContent) IsDir() bool { return false }

// ModTime implements fs.FileInfo.
func (r *fileInfoContent) ModTime() time.Time { panic("unsupported") }

// Mode implements fs.FileInfo.
func (r *fileInfoContent) Mode() fs.FileMode { panic("unsupported") }

// Name implements fs.FileInfo.
func (r *fileInfoContent) Name() string { return r.file.GetName() }

// Size implements fs.FileInfo.
func (r *fileInfoContent) Size() int64 { return int64(r.file.GetSize()) }

// Sys implements fs.FileInfo.
func (r *fileInfoContent) Sys() any { return r.file }

func NewFileInfo(contents *github.RepositoryContent) fs.FileInfo {
	return &fileInfoContent{contents}
}
