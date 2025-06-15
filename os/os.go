package os

import (
	"io"
	"io/fs"
	"os"
	"time"
)

const (
	O_RDONLY = os.O_RDONLY
	O_WRONLY = os.O_WRONLY
	O_RDWR   = os.O_RDWR
	O_APPEND = os.O_APPEND
	O_CREATE = os.O_CREATE
	O_EXCL   = os.O_EXCL
	O_SYNC   = os.O_SYNC
	O_TRUNC  = os.O_TRUNC

	DevNull = os.DevNull

	ModeAppend     = os.ModeAppend
	ModeCharDevice = os.ModeCharDevice
	ModeDevice     = os.ModeDevice
	ModeDir        = os.ModeDir
	ModeExclusive  = os.ModeExclusive
	ModeIrregular  = os.ModeIrregular
	ModeNamedPipe  = os.ModeNamedPipe
	ModePerm       = os.ModePerm
	ModeSetgid     = os.ModeSetgid
	ModeSetuid     = os.ModeSetuid
	ModeSocket     = os.ModeSocket
	ModeSticky     = os.ModeSticky
	ModeSymlink    = os.ModeSymlink
	ModeTemporary  = os.ModeTemporary
	ModeType       = os.ModeType
)

var (
	Args = os.Args

	ErrClosed           = os.ErrClosed
	ErrDeadlineExceeded = os.ErrDeadlineExceeded
	ErrExist            = os.ErrExist
	ErrInvalid          = os.ErrInvalid
	ErrNoDeadline       = os.ErrNoDeadline
	ErrNotExist         = os.ErrNotExist
	ErrPermission       = os.ErrPermission
	ErrProcessDone      = os.ErrProcessDone

	Interrupt = os.Interrupt

	NewFile         = os.NewFile
	NewSyscallError = os.NewSyscallError
)

type (
	DirEntry     = os.DirEntry
	File         = os.File
	FileInfo     = os.FileInfo
	FileMode     = os.FileMode
	LinkError    = os.LinkError
	PathError    = os.PathError
	ProcAttr     = os.ProcAttr
	Process      = os.Process
	ProcessState = os.ProcessState
	Root         = os.Root
	Signal       = os.Signal
	SyscallError = os.SyscallError
)

var System = sys{}

type sys struct{}

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

func (sys) Lchown(name string, uid, gid int) error {
	return os.Lchown(name, uid, gid)
}

func (sys) Link(oldname, newname string) error {
	return os.Link(oldname, newname)
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

func (sys) Stderr() io.Writer {
	return os.Stderr
}

func (sys) Stdin() io.Reader {
	return os.Stdin
}

func (sys) Stdout() io.Writer {
	return os.Stdout
}

func (sys) TempDir() string {
	return os.TempDir()
}
