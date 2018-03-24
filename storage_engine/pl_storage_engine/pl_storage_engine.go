package pl_storage_engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/moutansos/pinelakedb/storage_engine"
)

const masterFile = "db_config.json"

type PlStorageEngine struct {
	MasterFolder string                            `json:"-"` //Don't write to config file
	Queue        chan storage_engine.ExecutionStep `json:"-"`
}

// NewPlStorageEngine creates a new storage engine with a binary backend
func NewPlStorageEngine(MasterFolder string) (*PlStorageEngine, error) {
	se := new(PlStorageEngine)
	se.MasterFolder = MasterFolder
	se.Queue = make(chan storage_engine.ExecutionStep, 255)

	go se.StartQueueExecution()

	return se, nil
}

func (e *PlStorageEngine) Initialize() error {
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

func (e *PlStorageEngine) StartQueueExecution() {
	for s := range e.Queue {
		err := s.Execute()
		if err != nil {
			fmt.Println("error: " + err.Error())
			//TODO: Handle with a real logger
		}
	}
}

func (e *PlStorageEngine) AddStepToQueue(s storage_engine.ExecutionStep) {
	e.Queue <- s
}
