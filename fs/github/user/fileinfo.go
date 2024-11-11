package user

import (
	"io/fs"
	"os"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/unmango/go/fs/github/internal"
)

type FileInfo struct {
	internal.Fs
	User *github.User
}

// IsDir implements fs.FileInfo.
func (f *FileInfo) IsDir() bool {
	return true
}

// ModTime implements fs.FileInfo.
func (f *FileInfo) ModTime() time.Time {
	return f.User.GetUpdatedAt().Time
}

// Mode implements fs.FileInfo.
func (f *FileInfo) Mode() fs.FileMode {
	return os.ModeDir
}

// Name implements fs.FileInfo.
func (f *FileInfo) Name() string {
	return f.User.GetName()
}

// Size implements fs.FileInfo.
func (f *FileInfo) Size() int64 {
	panic("unimplemented")
}

// Sys implements fs.FileInfo.
func (f *FileInfo) Sys() any {
	return f.User
}
