package batch

import (
	"strings"
	"testing"
)

func TestRunCSV(t *testing.T) {
	input := strings.NewReader("operation,left,right\nadd,7,5\ndiv,10,2\n")

	results, err := RunCSV(input)
	if err != nil {
		t.Fatalf("RunCSV() error = %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("RunCSV() returned %d results, want 2", len(results))
	}
	if results[0].Result != 12 {
		t.Fatalf("first result = %v, want 12", results[0].Result)
	}
	if results[1].Result != 5 {
		t.Fatalf("second result = %v, want 5", results[1].Result)
	}
}

func TestRunCSVRejectsInvalidHeader(t *testing.T) {
	input := strings.NewReader("op,left,right\nadd,7,5\n")

	_, err := RunCSV(input)
	if err == nil {
		t.Fatal("RunCSV() error = nil, want invalid header error")
	}
}
