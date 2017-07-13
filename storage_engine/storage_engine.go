package storage_engine

type StorageEngine interface {
	Query()
	AddTable(name string)
}
