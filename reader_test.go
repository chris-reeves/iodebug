package iodebug

import (
	"log"
	"strings"
	"testing"
)

func TestReadReturnsExpectedData(t *testing.T) {
	// Configure logging to a string
	var logOutput strings.Builder
	log.SetOutput(&logOutput)

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

	// Test debug output - we must strip the log output prefix first as this
	// default to date so will vary for every test
	prefixLength := 20
	expectedOutput := [2]string{
		"00000000  54 68 69 73 20 69 73 20  61 20 74 65 73 74 20 6d  |This is a test m|",
		"00000010  65 73 73 61 67 65 21                              |essage!|",
	}
	receivedOutput := strings.Split(logOutput.String(), "\n")
	if len(receivedOutput) != len(expectedOutput) {
		t.Errorf("expected: len(receivedOutput) === %d, got: %d", len(expectedOutput), len(receivedOutput))
	}
	for i, _ := range expectedOutput {
		if receivedOutput[i][prefixLength:] != expectedOutput[i] {
			t.Errorf("expected: logOutput[%d] === %q, got: %q", i, expectedOutput[i][prefixLength:], receivedOutput[i])
		}
	}
}
