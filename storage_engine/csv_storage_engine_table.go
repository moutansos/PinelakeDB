package storage_engine

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
