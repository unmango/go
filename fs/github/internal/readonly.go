package internal

import "syscall"

type ReadOnlyFile struct{}

// Sync implements afero.File.
func (ReadOnlyFile) Sync() error {
	return syscall.EPERM
}

// Truncate implements afero.File.
func (ReadOnlyFile) Truncate(int64) error {
	return syscall.EPERM
}

// Write implements afero.File.
func (ReadOnlyFile) Write([]byte) (n int, err error) {
	return 0, syscall.EPERM
}

// WriteAt implements afero.File.
func (ReadOnlyFile) WriteAt([]byte, int64) (n int, err error) {
	return 0, syscall.EPERM
}

// WriteString implements afero.File.
func (ReadOnlyFile) WriteString(string) (ret int, err error) {
	return 0, syscall.EPERM
}
