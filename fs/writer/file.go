package writer

import (
	"io"
	"io/fs"
	"syscall"
)

type File struct {
	io.Writer
	name string
}

// Close implements afero.File.
func (w *File) Close() error {
	if c, ok := w.Writer.(io.WriteCloser); ok {
		return c.Close()
	}

	return nil
}

// Name implements afero.File.
func (w *File) Name() string {
	return w.name
}

// Read implements afero.File.
func (w *File) Read([]byte) (n int, err error) {
	return 0, syscall.EPERM
}

// ReadAt implements afero.File.
func (w *File) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, syscall.EPERM
}

// Readdir implements afero.File.
func (w *File) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, syscall.EPERM
}

// Readdirnames implements afero.File.
func (w *File) Readdirnames(n int) ([]string, error) {
	return nil, syscall.EPERM
}

// Seek implements afero.File.
func (w *File) Seek(offset int64, whence int) (int64, error) {
	return 0, syscall.EPERM
}

// Stat implements afero.File.
func (w *File) Stat() (fs.FileInfo, error) {
	return nil, syscall.EPERM
}

// Sync implements afero.File.
func (w *File) Sync() error {
	return syscall.EPERM
}

// Truncate implements afero.File.
func (w *File) Truncate(size int64) error {
	return syscall.EPERM
}

// Write implements afero.File.
// Subtle: this method shadows the method (Writer).Write of writerFile.Writer.
func (w *File) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}

// WriteAt implements afero.File.
func (w *File) WriteAt(p []byte, off int64) (n int, err error) {
	if wa, ok := w.Writer.(io.WriterAt); ok {
		return wa.WriteAt(p, off)
	}

	return 0, syscall.EPERM
}

// WriteString implements afero.File.
func (w *File) WriteString(s string) (ret int, err error) {
	return io.WriteString(w, s)
}
