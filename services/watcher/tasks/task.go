package tasks

type EventExecutor interface {
	GetEventData(data interface{}) ([]string, interface{}, error)
}
