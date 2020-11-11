package ws

import (
	"context"
	"encoding/json"

	"github.com/everstake/teztracker/ws/models"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

const rwBufferSize = 1024
const countOfLastMessages = 10

type (
	HubInterface interface {
		RegisterClient(*websocket.Conn)
		Broadcast(...interface{}) error
		BroadcastPrivate(uuid.UUID, ...interface{}) error
	}

	// Hub maintains the set of active clients and broadcasts messages to clients
	Hub struct {
		// All clients
		clients map[*Client]bool

		// Messages to all connections
		broadcast chan *PublicMsg

		// Register requests from the clients.
		registerChan chan *Client

		// Unregister requests from clients.
		unregisterChan chan *Client

		// connection upgrader
		upgrader *websocket.Upgrader
	}

	PublicMsg struct {
		channel string
		data    []byte
	}

	PrivateMsg struct {
		userUuid *uuid.UUID
		data     []byte
	}
)

func NewHub() *Hub {

	return &Hub{
		broadcast:      make(chan *PublicMsg),
		registerChan:   make(chan *Client),
		unregisterChan: make(chan *Client),

		clients: make(map[*Client]bool),

		upgrader: &websocket.Upgrader{
			ReadBufferSize:  rwBufferSize,
			WriteBufferSize: rwBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

// Title returns the title.
func (this *Hub) Title() string {
	return "Websocket Hub"
}

func (this *Hub) Run() {
	for {
		select {
		case client := <-this.registerChan:
			this.addClient(client)
		case client := <-this.unregisterChan:
			this.removeClient(client)
		case message := <-this.broadcast:
			log.Debug("Public broadcast", "clients", len(this.clients), "channel", message.channel)
			for c, _ := range this.clients {
				// Check if client has subscribed to this channel
				if c.isSubscribed(message.channel) {
					this.sendToClient(c, message.data)
				}
			}
		}
	}
}

// TODO GracefulStop shuts down the socket.
func (h *Hub) GracefulStop(ctx context.Context) error {
	//TODO: cancel all open amqp channels
	return nil
}

func (this *Hub) GetUpgrader() *websocket.Upgrader {
	return this.upgrader
}

func (this *Hub) sendToClient(c *Client, data []byte) {
	select {
	case c.sendChan <- data:
	default:
		this.removeClient(c)
	}
}

func (this *Hub) addClient(c *Client) {
	if _, ok := this.clients[c]; !ok {
		this.clients[c] = true
	}
}

func (this *Hub) removeClient(c *Client) {
	if _, ok := this.clients[c]; ok {
		delete(this.clients, c)
	}
}

func (this *Hub) RegisterClient(conn *websocket.Conn) {
	log.Debug("Incoming client connection", "address", conn.RemoteAddr().String())
	NewClient(this, conn)
}

func (this *Hub) Broadcast(msg models.PublicMessageInterface) error {
	data, err := this.serializeMessage(msg)
	if err != nil {
		return err
	}

	this.broadcast <- &PublicMsg{
		channel: msg.GetChannel(),
		data:    data,
	}

	return nil
}

func (this *Hub) serializeMessage(msg models.MessageInterface) ([]byte, error) {
	bm := models.BasicMessage{
		Event: msg.GetEvent(),
		Data:  msg,
	}

	return json.Marshal(bm)
}
