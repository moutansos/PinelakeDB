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
