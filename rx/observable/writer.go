package observable

import (
	"io"

	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/subject"
)

type Writer interface {
	io.WriteCloser
	rx.Observable[[]byte]
}

type writer struct {
	rx.Subject[[]byte]
}

// Close implements io.WriteCloser.
func (w *writer) Close() error {
	w.OnComplete()
	return nil
}

// Write implements io.WriteCloser.
func (w *writer) Write(p []byte) (n int, err error) {
	w.OnNext(p)
	return len(p), nil
}

func NewWriter() Writer {
	return &writer{subject.New[[]byte]()}
}
