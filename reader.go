package iodebug

import (
	"io"
)

type DebugReader struct {
	subReader io.Reader
}

func NewDebugReader(r io.Reader) io.Reader {
	reader := new(DebugReader)
	reader.subReader = r
	return reader
}

func (r *DebugReader) Read(p []byte) (n int, err error) {
	n, err = r.subReader.Read(p)
	return n, err
}
