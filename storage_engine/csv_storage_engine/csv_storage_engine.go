package csv_storage_engine

import (
	"encoding/json"
	"io/ioutil"
)

type CsvStorageEngine struct {
	MasterFilePath string                   `json:"-"` //Don't write this field to file
	Tables         []*CsvStorageEngineTable `json:"tables"`
}

// NewTable adds a new table to a CsvStoarageEngine
func (se *CsvStorageEngine) NewTable(name string) error {
	//TODO: Map file path
	table, err := NewCsvStorageEngineTable(name)
	if err != nil {
		return err
	}
	se.Tables = append(se.Tables, table)
	return nil
}

// NewCsvStorageEngine creates a new storage engine with a csv backend at the specified path
func NewCsvStorageEngine(pathOfMaster string) (*CsvStorageEngine, error) {
	var se CsvStorageEngine
	se.MasterFilePath = pathOfMaster
	return &se, nil
}

func writeEngineMasterFile(e *CsvStorageEngine) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(e.MasterFilePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
