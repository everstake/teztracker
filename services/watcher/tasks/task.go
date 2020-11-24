package tasks

import wsmodels "github.com/everstake/teztracker/ws/models"

type EventExecutor interface {
	GetEventData(data interface{}) ([]wsmodels.EventType, interface{}, error)
}
