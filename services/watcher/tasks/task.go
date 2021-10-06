package tasks

type EventExecutor interface {
	GetEventData(data interface{}) (interface{}, error)
}
