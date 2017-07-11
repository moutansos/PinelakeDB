package storage_engine

import (
	"errors"
)

type CsvStorageEngineTable struct {
	ID         string
	FilePath   string
	fieldNames []string
}

// AddField adds a new field or column to the csv storage engine table
func (t *CsvStorageEngineTable) AddField(name string) error {
	//TODO: name validation
	//TODO: shift all columns to accomidate
	t.fieldNames = append(t.fieldNames, name)
	return nil
}

// RemoveField removes the field or column from the csv storage engine table
func (t *CsvStorageEngineTable) RemoveField(name string) error {
	ord, err := t.GetOrdinal(name)
	if err != nil {
		return err
	}

	t.fieldNames = append(t.fieldNames[:*ord], t.fieldNames[*ord+1:]...)
	return nil
}

// GetOrdinal gets the index of the field or column with the given name
func (t *CsvStorageEngineTable) GetOrdinal(name string) (*int, error) {
	for i, field := range t.fieldNames {
		if field == name {
			return &i, nil
		}
	}
	return nil, errors.New("Invalid field name")
}
