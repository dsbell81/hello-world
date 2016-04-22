package calc

import "testing"

func TestAdd(t *testing.T) {
	var result int
	result = Add(1, 2)
	if result != 3 {
		t.Error("Expected 3, got: ", result)
	}
}

func TestSubtract(t *testing.T) {
	var result int
	result = Subtract(9, 5)
	if result != 4 {
		t.Error("Expected 4, got ", result)
	}
}
