package writer

import (
	"io"
	"io/fs"
	"syscall"
)

type File struct {
	w    io.Writer
	name string
}

// Close implements afero.File.
func (f *File) Close() error {
	return nil
}

// Name implements afero.File.
func (f *File) Name() string {
	return f.name
}

// Read implements afero.File.
func (f *File) Read([]byte) (n int, err error) {
	return 0, syscall.EPERM
}

// ReadAt implements afero.File.
func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (f *File) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, syscall.EPERM
}

// Readdirnames implements afero.File.
func (f *File) Readdirnames(n int) ([]string, error) {
	return nil, syscall.EPERM
}

// Seek implements afero.File.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (f *File) Stat() (fs.FileInfo, error) {
	return nil, syscall.EPERM
}

// Sync implements afero.File.
func (f *File) Sync() error {
	return syscall.EPERM
}

// Truncate implements afero.File.
func (f *File) Truncate(size int64) error {
	return syscall.EPERM
}

// Write implements afero.File.
// Subtle: this method shadows the method (Writer).Write of writerFile.Writer.
func (f *File) Write(p []byte) (n int, err error) {
	return f.w.Write(p)
}

// WriteAt implements afero.File.
func (f *File) WriteAt(p []byte, off int64) (n int, err error) {
	if wa, ok := f.w.(io.WriterAt); ok {
		return wa.WriteAt(p, off)
	}

	return 0, syscall.EPERM
}

// WriteString implements afero.File.
func (f *File) WriteString(s string) (ret int, err error) {
	return io.WriteString(f.w, s)
}
