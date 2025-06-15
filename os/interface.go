package os

import (
	"io"
	"io/fs"
	"time"
)

// TODO: Implement the rest of the os package API

type Os interface {
	Chdir(dir string) error
	Chmod(name string, mode FileMode) error
	Chown(name string, uid, gid int) error
	Chtimes(name string, atime, mtime time.Time) error
	Clearenv()
	CopyFS(dir string, fsys fs.FS) error
	DirFS(dir string) fs.FS
	Environ() []string
	Executable() (string, error)
	Exit(code int)
	Expand(s string, mapping func(string) string) string
	ExpandEnv(s string) string
	Getegid() int
	Getenv(key string) string
	Geteuid() int
	Getgid() int
	Getgroups() ([]int, error)
	Getpagesize() int
	Getpid() int
	Getppid() int
	Getuid() int
	Getwd() (string, error)
	Hostname() (name string, err error)
	Lchown(name string, uid, gid int) error
	Link(oldname, newname string) error
	LookupEnv(key string) (string, bool)
	Mkdir(name string, mode FileMode) error
	MkdirAll(path string, mode FileMode) error
	MkdirTemp(dir, pattern string) (string, error)
	Stderr() io.Writer
	Stdin() io.Reader
	Stdout() io.Writer
	TempDir() string
}
