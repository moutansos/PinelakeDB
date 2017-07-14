package pl_storage_engine

import (
	"os"
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
