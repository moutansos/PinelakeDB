package storage_engine

import (
	"errors"
	"path"
)

type CsvStorageEngineTable struct {
	ID         string   `json:"id"`
	FilePath   string   `json:"filePath"`
	FieldNames []string `json:"fieldNames"`
}

// NewCsvStorageEngineTable returns a new csv storage engine table
func NewCsvStorageEngineTable(name string) (*CsvStorageEngineTable, error) {
	//TODO: Validate name
	var newTable CsvStorageEngineTable
	newTable.ID = name
	//newTable.FilePath = csvFile
	return &newTable, nil
}

// AddField adds a new field or column to the csv storage engine table
func (t *CsvStorageEngineTable) AddField(name string) error {
	//TODO: name validation
	//TODO: shift all columns to accomidate
	t.FieldNames = append(t.FieldNames, name)
	return nil
}

// RemoveField removes the field or column from the csv storage engine table
func (t *CsvStorageEngineTable) RemoveField(name string) error {
	ord, err := t.GetOrdinal(name)
	if err != nil {
		return err
	}

	t.FieldNames = append(t.FieldNames[:*ord], t.FieldNames[*ord+1:]...)
	return nil
}

// GetOrdinal gets the index of the field or column with the given name
func (t *CsvStorageEngineTable) GetOrdinal(name string) (*int, error) {
	for i, field := range t.FieldNames {
		if field == name {
			return &i, nil
		}
	}
	return nil, errors.New("Invalid field name")
}

func (t *CsvStorageEngineTable) genTableFileName(mainFolderPath string) error {
	dir := path.Join(mainFolderPath, t.ID) //TODO: Add csv extension
	t.FilePath = dir
	return nil
}
