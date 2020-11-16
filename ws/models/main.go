package models

type EventType string

const (
	eventTypeSystem    EventType = "sys"
	eventTypeBlock     EventType = "block"
	eventTypeOperation EventType = "operation"
)

type sysMessage string

const SysMessageHello sysMessage = "hello"
const SysMessageUnknownCommand sysMessage = "unknown_command"
const SysMessageSubscribed sysMessage = "subscribed"
const SysMessageUnsubscribed sysMessage = "unsubscribed"

type BasicMessage struct {
	Event EventType   `json:"event"`
	Data  interface{} `json:"data"`
}

func (bm BasicMessage) GetEvent() EventType {
	return bm.Event
}

type SystemMessage struct {
	Message     sysMessage `json:"msg"`
	Description string     `json:"description"`
}

type MessageInterface interface {
	GetEvent() EventType
}

func (sm SystemMessage) GetEvent() EventType {
	return eventTypeSystem
}

type PublicMessageInterface interface {
	GetChannel() string
	GetEvent() EventType
}
