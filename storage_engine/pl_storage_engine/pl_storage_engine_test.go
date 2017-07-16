package pl_storage_engine

import (
	"os"
	"path"
	"testing"
)

func TestWriteEngineMasterFile(t *testing.T) {
	dir := os.TempDir()
	eng := PlStorageEngine{
		MasterFolder: dir,
	}
	err := writeEngineMasterFile(&eng)
	if err != nil {
		t.Error(err)
	}
	//TODO: Check to make sure the file contains what it should
}

func TestInsert(t *testing.T) {
	dir := os.TempDir()
	fPath := path.Join(dir, "testing.bin")
	f, err := os.Create(fPath)
	if err != nil {
		t.Error(err)
	}

	se, err := NewPlStorageEngine(dir)
	if err != nil {
		t.Error(err)
	}

	row := PlStorageEngineRow{
		Index:   1,
		Pointer: 65,
	}

	n, err := se.Insert(f, row)
	if err != nil {
		t.Error(err)
	} else if n != 16 {
		t.Errorf("Number of bytes written was wrong. Was only %v bytes", n)
	}
}
