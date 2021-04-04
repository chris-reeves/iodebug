package iodebug

import (
	"strings"
	"testing"
)

func TestReadReturnsExpectedData(t *testing.T) {
	// Set up a sub-reader
	str := "This is a test!"
	strReader := strings.NewReader(str)

	// Create a DebugReader and then Read()
	r := NewDebugReader(strReader)
	buff := make([]byte, 20)
	n, err := r.Read(buff)

	// Test data returned by Read()
	if n != 15 {
		t.Errorf("expected: n === 15, got: %d", n)
	}
	if err != nil {
		t.Errorf("expected: err === nil, got: %q", err)
	}
	if string(buff[:n]) != str {
		t.Errorf("expected: buff[:n] === %q, got: %q", str, string(buff[:n]))
	}
}
