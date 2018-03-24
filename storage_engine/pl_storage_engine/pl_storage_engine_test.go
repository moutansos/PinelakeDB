package pl_storage_engine

import (
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/oklog/ulid"
)

func generateTestDirectory() (dir string, err error) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	dir = path.Join(os.TempDir(), id.String())
	err = os.Mkdir(dir, 0644)
	if err != nil {
		return "", err
	}
	return dir, nil
}
func TestWriteEngineMasterFile(t *testing.T) {
	dir, err := generateTestDirectory()
	if err != nil {
		t.Error(err)
	}
	eng, err := NewPlStorageEngine(dir)
	if err != nil {
		t.Error(err)
	}
	err = eng.writeEngineMasterFile()
	if err != nil {
		t.Error(err)
	}
	//TODO: Check to make sure the file contains what it should
}
