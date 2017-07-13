package storage_engine

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteEngineMasterFile(t *testing.T) {
	dir := os.TempDir()
	dir = filepath.Join(dir, "testMasterFile.json")
	eng := CsvStorageEngine{
		MasterFilePath: dir,
	}
	err := writeEngineMasterFile(&eng)
	if err != nil {
		t.Error(err)
	}
}

func TestNewCsvStorageEngine(t *testing.T) {
	dir := os.TempDir()
	dir = filepath.Join(dir, "testMasterFile2.json")
	se, err := NewCsvStorageEngine(dir)
	if err != nil {
		t.Error(err)
	} else if se.MasterFilePath != dir {
		t.Error("Wrong master file path")
	}
}

func TestNewTable(t *testing.T) {
	dir := os.TempDir()
	dir = filepath.Join(dir, "testMasterFile.json")
	eng := CsvStorageEngine{
		MasterFilePath: dir,
	}
	eng.NewTable("Test3")
	if len(eng.Tables) != 1 {
		t.Error("There was the wrong number of tables")
	} else if eng.Tables[0].ID != "Test3" {
		t.Error("There was an invalid first table added to the engine")
	}
}
