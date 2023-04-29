package functions

import (
	"testing"
)

func TestIsEqual(t *testing.T) {

	res := Sum(1, 2)

	if res != 3 {
		t.Errorf("Expected 3, got %d", res)
	}
}

func TestNotIsEqual(t *testing.T) {

	res := Sum(3, 2)

	if res != 5 {
		t.Errorf("Expected 3, got %d", res)
	}
}

func TestIsPalatial(t *testing.T) {

	if !IsPalatial("ana") {
		t.Errorf("Expected true, got false")
	}
}

func TestNotPalatial(t *testing.T) {

	if IsPalatial("icaro") {
		t.Errorf("Expected false, got true")
	}

}
