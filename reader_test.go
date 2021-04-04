package iodebug

import (
	"strings"
	"testing"
)

func TestReadReturnsExpectedData(t *testing.T) {
	// Set up a sub-reader
	str := "This is a test message!"
	strReader := strings.NewReader(str)

	// Create a DebugReader and then Read()
	r := NewDebugReader(strReader)
	buff := make([]byte, 32)
	n, err := r.Read(buff)

	// Test data returned by Read()
	if n != len(str) {
		t.Errorf("expected: n === %d, got: %d", len(str), n)
	}
	if err != nil {
		t.Errorf("expected: err === nil, got: %q", err)
	}
	if string(buff[:n]) != str {
		t.Errorf("expected: buff[:n] === %q, got: %q", str, string(buff[:n]))
	}
}
