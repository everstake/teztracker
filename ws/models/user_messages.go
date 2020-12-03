package models

const ClientMessageTypeSubscribe string = "subscribe"
const ClientMessageTypeUnsubscribe string = "unsubscribe"

type ClientMessage interface{}

type ClientMessageSubscribe struct {
	Channels []string
}

type ClientMessageUnsubscribe struct {
	Channels []string
}
