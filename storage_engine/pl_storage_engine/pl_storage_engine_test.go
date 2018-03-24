package pl_storage_engine

import (
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/moutansos/pinelakedb/storage_engine"

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
func TestWriteEngineMasterFileSync(t *testing.T) {
	dir, err := generateTestDirectory()
	if err != nil {
		t.Error(err)
	}
	eng, err := NewPlStorageEngine(dir)
	if err != nil {
		t.Error(err)
	}
	err = eng.Initialize()
	if err != nil {
		t.Error(err)
	}
	//TODO: Check to make sure the file contains what it should
}

func TestWriteEngineMasterFileAsync(t *testing.T) {
	dir, err := generateTestDirectory()
	if err != nil {
		t.Error(err)
	}
	eng, err := NewPlStorageEngine(dir)
	if err != nil {
		t.Error(err)
	}

	//Create an initialization execution step
	step := storage_engine.NewInitializationExecutionStep(eng)
	eng.AddStepToQueue(step)
	//TODO: Confirm results
}
