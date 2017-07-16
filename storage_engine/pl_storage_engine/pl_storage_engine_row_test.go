package pl_storage_engine

import "testing"

func TestIsValidRowIndexAddr(t *testing.T) {
	valid := []uint64{0, 7, 15, 23, 31}
	invalid := []uint64{67, 43, 342, 67, 98}

	for _, n := range valid {
		if !IsValidRowIndexAddr(n) {
			t.Errorf("Valid index of %v was verified to be invalid.", n)
		}
	}

	for _, n := range invalid {
		if IsValidRowIndexAddr(n) {
			t.Errorf("Invalid index of %v was verified to be valid.", n)
		}
	}
}

func TestIsValidRowPointerAddr(t *testing.T) {
	valid := []uint64{8, 16, 24, 32}
	invalid := []uint64{67, 43, 342, 62, 98}

	for _, n := range valid {
		if !IsValidRowPointerAddr(n) {
			t.Errorf("Valid pointer of %v was verified to be invalid.", n)
		}
	}

	for _, n := range invalid {
		if IsValidRowPointerAddr(n) {
			t.Errorf("Invalid pointer of %v was verified to be valid.", n)
		}
	}
}
