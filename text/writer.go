package text

import "io"

// This probably already exists somewhere I just couldn't find it

type delimWriter struct {
	writer io.Writer
	delim  rune
}

// Write implements io.Writer.
func (d *delimWriter) Write(p []byte) (n int, err error) {
	return d.writer.Write(append(p, byte(d.delim)))
}

func NewDelimWriter(w io.Writer, delim rune) io.Writer {
	return &delimWriter{w, delim}
}
