package units

import "testing"

func TestConvertLength(t *testing.T) {
	got, err := ConvertLength(1200, "meter", "kilometer")
	if err != nil {
		t.Fatalf("ConvertLength() error = %v", err)
	}
	if got != 1.2 {
		t.Fatalf("ConvertLength() = %v, want 1.2", got)
	}
}

func TestConvertLengthUnsupportedUnit(t *testing.T) {
	_, err := ConvertLength(1, "parsec", "meter")
	if err == nil {
		t.Fatal("ConvertLength() error = nil, want unsupported unit error")
	}
}
