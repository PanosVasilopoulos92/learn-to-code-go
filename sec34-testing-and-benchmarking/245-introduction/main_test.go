package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 21
	result := add(1, 2, 3, 4, 5, 6)
	if expected != result {
		t.Errorf("Expected: %d, got: %d", expected, result)
	}
}
