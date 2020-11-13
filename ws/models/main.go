package models

type eventType string

const (
	eventTypeSystem    eventType = "sys"
	eventTypeBlock     eventType = "block"
	eventTypeOperation eventType = "operation"
)

type sysMessage string

const SysMessageHello sysMessage = "hello"
const SysMessageUnknownCommand sysMessage = "unknown_command"
const SysMessageSubscribed sysMessage = "subscribed"
const SysMessageUnsubscribed sysMessage = "unsubscribed"

type BasicMessage struct {
	Event eventType   `json:"event"`
	Data  interface{} `json:"data"`
}

func (bm BasicMessage) GetEvent() eventType {
	return bm.Event
}

type SystemMessage struct {
	Message     sysMessage `json:"msg"`
	Description string     `json:"description"`
}

type MessageInterface interface {
	GetEvent() eventType
}

func (sm SystemMessage) GetEvent() eventType {
	return eventTypeSystem
}

type PublicMessageInterface interface {
	GetChannel() string
	GetEvent() eventType
}
