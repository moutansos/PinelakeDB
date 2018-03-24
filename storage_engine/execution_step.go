package storage_engine

type ExecutionStep interface {
	Execute() (err error)
}
