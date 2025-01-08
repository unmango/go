package make

import (
	"bufio"
	"io"
)

type Writer struct {
	w *bufio.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{bufio.NewWriter(w)}
}

func (w *Writer) WriteLine() (n int, err error) {
	return w.w.WriteRune('\n')
}

func (w *Writer) WriteTarget(t string) (n int, err error) {
	return io.WriteString(w.w, t+":")
}

func (w *Writer) WriteTargets(ts []string) (n int, err error) {
	for i, t := range ts {
		if i+1 < len(ts) {
			t = t + " "
		}
		if x, err := io.WriteString(w.w, t); err != nil {
			return 0, err
		} else {
			n += x
		}
	}

	if x, err := io.WriteString(w.w, ":"); err != nil {
		return 0, err
	} else {
		return n + x, nil
	}
}
