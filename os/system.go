package os

import (
	"io"
	"io/fs"
	"os"
	"time"
)

var System = sys{}

type sys struct{}

func (sys) Args() []string {
	return os.Args
}

func (sys) Chdir(dir string) error {
	return os.Chdir(dir)
}

func (sys) Chmod(name string, mode FileMode) error {
	return os.Chmod(name, mode)
}

func (sys) Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

func (sys) Chtimes(name string, atime, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func (sys) Clearenv() {
	os.Clearenv()
}

func (sys) CopyFS(dir string, fsys fs.FS) error {
	return os.CopyFS(dir, fsys)
}

func (sys) DirFS(dir string) fs.FS {
	return os.DirFS(dir)
}

func (sys) Environ() []string {
	return os.Environ()
}

func (sys) Executable() (string, error) {
	return os.Executable()
}

func (sys) Exit(code int) {
	os.Exit(code)
}

func (sys) Expand(s string, mapping func(string) string) string {
	return os.Expand(s, mapping)
}

func (sys) ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func (sys) Getegid() int {
	return os.Getegid()
}

func (sys) Getenv(key string) string {
	return os.Getenv(key)
}

func (sys) Geteuid() int {
	return os.Geteuid()
}

func (sys) Getgid() int {
	return os.Getgid()
}

func (sys) Getgroups() ([]int, error) {
	return os.Getgroups()
}

func (sys) Getpagesize() int {
	return os.Getpagesize()
}

func (sys) Getpid() int {
	return os.Getpid()
}

func (sys) Getppid() int {
	return os.Getppid()
}

func (sys) Getuid() int {
	return os.Getuid()
}

func (sys) Getwd() (dir string, err error) {
	return os.Getwd()
}

func (sys) Hostname() (name string, err error) {
	return os.Hostname()
}

func (sys) IsPathSeparator(c uint8) bool {
	return os.IsPathSeparator(c)
}

func (sys) Lchown(name string, uid, gid int) error {
	return os.Lchown(name, uid, gid)
}

func (sys) Link(oldname, newname string) error {
	return os.Link(oldname, newname)
}

func (sys) Lstat(name string) (FileInfo, error) {
	return os.Lstat(name)
}

func (sys) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (sys) Mkdir(name string, perm FileMode) error {
	return os.Mkdir(name, perm)
}

func (sys) MkdirAll(path string, perm FileMode) error {
	return os.MkdirAll(path, perm)
}

func (sys) MkdirTemp(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

func (sys) OpenInRoot(dir, name string) (*File, error) {
	return os.OpenInRoot(dir, name)
}

func (sys) OpenRoot(name string) (*Root, error) {
	return os.OpenRoot(name)
}

func (sys) Pipe() (r io.Reader, w io.Writer, err error) {
	return os.Pipe()
}

func (sys) ReadDir(name string) ([]DirEntry, error) {
	return os.ReadDir(name)
}

func (sys) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (sys) Readlink(name string) (string, error) {
	return os.Readlink(name)
}

func (sys) Remove(name string) error {
	return os.Remove(name)
}

func (sys) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (sys) Rename(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (sys) SameFile(fi1 FileInfo, fi2 FileInfo) bool {
	return os.SameFile(fi1, fi2)
}

func (sys) Setenv(key string, value string) error {
	return os.Setenv(key, value)
}

func (sys) Stat(name string) (FileInfo, error) {
	return os.Stat(name)
}

func (sys) Stderr() io.Writer {
	return os.Stderr
}

func (sys) Stdin() io.Reader {
	return os.Stdin
}

func (sys) Stdout() io.Writer {
	return os.Stdout
}

func (sys) Symlink(oldname string, newname string) error {
	return os.Symlink(oldname, newname)
}

func (sys) TempDir() string {
	return os.TempDir()
}

func (sys) Truncate(name string, size int64) error {
	return os.Truncate(name, size)
}

func (sys) Unsetenv(key string) error {
	return os.Unsetenv(key)
}

func (sys) UserCacheDir() (string, error) {
	return os.UserCacheDir()
}

func (sys) UserConfigDir() (string, error) {
	return os.UserConfigDir()
}

func (sys) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (sys) WriteFile(name string, data []byte, perm FileMode) error {
	return os.WriteFile(name, data, perm)
}
