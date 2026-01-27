package os

import "os"

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
	Args      = os.Args
	Stderr    = os.Stderr
	Stdin     = os.Stdin
	Stdout    = os.Stdout
	Interrupt = os.Interrupt

	ErrClosed           = os.ErrClosed
	ErrDeadlineExceeded = os.ErrDeadlineExceeded
	ErrExist            = os.ErrExist
	ErrInvalid          = os.ErrInvalid
	ErrNoDeadline       = os.ErrNoDeadline
	ErrNotExist         = os.ErrNotExist
	ErrPermission       = os.ErrPermission
	ErrProcessDone      = os.ErrProcessDone
	NewSyscallError     = os.NewSyscallError

	Chdir       = os.Chdir
	Clearenv    = os.Clearenv
	Environ     = os.Environ
	Executable  = os.Executable
	Exit        = os.Exit
	Expand      = os.Expand
	ExpandEnv   = os.ExpandEnv
	Getenv      = os.Getenv
	Getpagesize = os.Getpagesize
	Getegid     = os.Getegid
	Geteuid     = os.Geteuid
	Getgid      = os.Getgid
	Getuid      = os.Getuid
	LookupEnv   = os.LookupEnv
	Setenv      = os.Setenv
	Unsetenv    = os.Unsetenv
	UserHomeDir = os.UserHomeDir

	Chmod      = os.Chmod
	Chown      = os.Chown
	Chtimes    = os.Chtimes
	Create     = os.Create
	CreateTemp = os.CreateTemp
	DirFS      = os.DirFS
	CopyFS     = os.CopyFS
	Open       = os.Open
	OpenFile   = os.OpenFile
	OpenInRoot = os.OpenInRoot
	OpenRoot   = os.OpenRoot
	Lchown     = os.Lchown
	Link       = os.Link
	Mkdir      = os.Mkdir
	MkdirAll   = os.MkdirAll
	MkdirTemp  = os.MkdirTemp
	NewFile    = os.NewFile
	Remove     = os.Remove
	RemoveAll  = os.RemoveAll
	Rename     = os.Rename
	Stat       = os.Stat
	Symlink    = os.Symlink
	Truncate   = os.Truncate
	WriteFile  = os.WriteFile
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
