package constants

import "testing"

func TestLookup(t *testing.T) {
	value, exists := Lookup("PI")
	if !exists {
		t.Fatal("Lookup() did not find PI")
	}
	if value < 3.14 || value > 3.15 {
		t.Fatalf("Lookup() = %v, want approximately pi", value)
	}
}

func TestNames(t *testing.T) {
	names := Names()
	if len(names) != 3 {
		t.Fatalf("Names() returned %d names, want 3", len(names))
	}
	if names[0] != "e" || names[1] != "pi" || names[2] != "tau" {
		t.Fatalf("Names() = %v, want sorted names", names)
	}
}
