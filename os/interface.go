package os

import (
	"io"
	"io/fs"
	"time"
)

// TODO: Figure out what to do with *File and the io.* types

type Ent interface {
	Args() []string
	Executable() (string, error)
	Exit(code int)
}

type Env interface {
	Clearenv()
	Environ() []string
	ExpandEnv(s string) string
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
	Setenv(key, value string) error
	Unsetenv(key string) error
	UserCacheDir() (string, error)
	UserConfigDir() (string, error)
	UserHomeDir() (string, error)
}

type Fs interface {
	CopyFS(dir string, fsys fs.FS) error
	DirFS(dir string) fs.FS
	Chdir(dir string) error
	Chmod(name string, mode FileMode) error
	Chown(name string, uid, gid int) error
	Chtimes(name string, atime, mtime time.Time) error
	Getwd() (string, error)
	IsPathSeparator(c uint8) bool
	Lchown(name string, uid, gid int) error
	Link(oldname, newname string) error
	Mkdir(name string, mode FileMode) error
	MkdirAll(path string, mode FileMode) error
	MkdirTemp(dir, pattern string) (string, error)
	Pipe() (r io.Reader, w io.Writer, err error)
	ReadDir(name string) ([]DirEntry, error)
	ReadFile(name string) ([]byte, error)
	Readlink(name string) (string, error)
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldpath, newpath string) error
	SameFile(fi1, fi2 FileInfo) bool
	Symlink(oldname, newname string) error
	TempDir() string
	Truncate(name string, size int64) error
	WriteFile(name string, data []byte, perm FileMode) error
}

type Id interface {
	Getegid() int
	Geteuid() int
	Getgid() int
	Getgroups() ([]int, error)
	Getuid() int
}

type Net interface {
	Hostname() (name string, err error)
}

type Stdio interface {
	Stderr() io.Writer
	Stdin() io.Reader
	Stdout() io.Writer
}

type Sys interface {
	Getpagesize() int
	Getpid() int
	Getppid() int
}

type Os interface {
	Ent
	Env
	Fs
	Id
	Net
	Stdio
	Sys
}
