package pl_storage_engine

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

const masterFile = "db_config.json"

type PlStorageEngine struct {
	MasterFolder string `json:"-"` //Don't write to config file
}

// NewPlStorageEngine creates a new storage engine with a binary backend
func NewPlStorageEngine(MasterFolder string) (*PlStorageEngine, error) {
	se := PlStorageEngine{
		MasterFolder: MasterFolder,
	}
	//TODO: Generate a config file and write it out
	return &se, nil
}

func writeEngineMasterFile(e *PlStorageEngine) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	masterFileName := path.Join(e.MasterFolder, masterFile)
	err = ioutil.WriteFile(masterFileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
