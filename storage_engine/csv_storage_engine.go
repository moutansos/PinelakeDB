package storage_engine

type CsvStorageEngine struct {
	MasterFilePath string
}

// NewCsvStorageEngine creates a new storage engine with a csv backend at the specified path
func NewCsvStorageEngine(pathOfMaster string) (*CsvStorageEngine, error) {
	var se CsvStorageEngine
	se.MasterFilePath = pathOfMaster
	return &se, nil
}
