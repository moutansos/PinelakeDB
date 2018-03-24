package storage_engine

type InitializationExecutionStep struct {
	Engine StorageEngine
}

func NewInitializationExecutionStep(e StorageEngine) *InitializationExecutionStep {
	step := new(InitializationExecutionStep)
	step.Engine = e
	return step
}

func (s *InitializationExecutionStep) Execute() (err error) {
	return s.Engine.Initialize()
}
