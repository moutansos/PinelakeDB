package pl_storage_engine

import (
	"os"
	"testing"
)

func TestIsValidRowIndexAddr(t *testing.T) {
	valid := []uint64{0, 16, 32, 48, 64}
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
	valid := []uint64{8, 24, 40, 56, 72}
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

func TestWriteIndex(t *testing.T) {
	//dir := os.TempDir()
	//fPath := path.Join(dir, "testing.bin")
	//f, err := os.Create(fPath)
	f, err := os.Create("testing.bin")
	if err != nil {
		t.Error(err)
	}

	row := PlStorageEngineRow{
		Index:   1,
		Pointer: 65,
	}

	n, err := writeIndex(f, 0, row)
	if err != nil {
		t.Error(err)
	} else if n != 16 {
		t.Errorf("Number of bytes written was wrong. Was only %v bytes", n)
	}

	row2 := PlStorageEngineRow{
		Index:   2,
		Pointer: 76,
	}

	n2, err := writeIndex(f, 16, row2)
	if err != nil {
		t.Error(err)
	} else if n2 != 16 {
		t.Errorf("Number of bytes written was wrong. Was only %v bytes", n)
	}
}
