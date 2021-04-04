package iodebug

import (
	"bufio"
	"encoding/hex"
	"io"
	"log"
	"strings"
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

	// Hex dump of data to logger, line by line
	scanner := bufio.NewScanner(strings.NewReader(hex.Dump(p[:n])))
	for scanner.Scan() {
		log.Print(scanner.Text())
	}

	return n, err
}
