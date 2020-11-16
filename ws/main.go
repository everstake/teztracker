package ws

import (
	"context"
	"encoding/json"

	"github.com/everstake/teztracker/ws/models"
	"github.com/gorilla/websocket"
	"github.com/mailru/easygo/netpoll"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const rwBufferSize = 1024
const countOfLastMessages = 10

type (
	HubInterface interface {
		RegisterClient(*websocket.Conn)
		Broadcast(...interface{}) error
	}

	// Hub maintains the set of active clients and broadcasts messages to clients
	Hub struct {
		//Context
		ctx context.Context

		//Cancel func
		cancel context.CancelFunc

		//
		poller netpoll.Poller

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
)

func NewHub() *Hub {

	ctx, cancel := context.WithCancel(context.Background())
	poller, err := netpoll.New(nil)
	if err != nil {

	}

	return &Hub{
		ctx:    ctx,
		cancel: cancel,
		poller: poller,

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
func (h *Hub) Title() string {
	return "Websocket Hub"
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.registerChan:

			// Make conn to be observed by netpoll instance.
			// Note that EventRead is identical to EPOLLIN on Linux.
			h.poller.Start(client.desc, func(event netpoll.Event) {
				// We spawn goroutine here to prevent poller wait loop
				// to become locked during receiving packet from ch.
				go client.Reader()
			})

			h.addClient(client)

		case client := <-h.unregisterChan:
			h.removeClient(client)

			h.poller.Stop(client.desc)
		case message := <-h.broadcast:
			log.Debug("Public broadcast", "clients", len(h.clients), "channel", message.channel)
			for c, _ := range h.clients {
				// Check if client has subscribed to this channel
				if c.isSubscribed(message.channel) {
					h.sendToClient(c, message.data)
				}
			}
		}
	}
}

// TODO GracefulStop shuts down the socket.
func (h *Hub) GracefulStop(ctx context.Context) error {

	h.cancel()
	return nil
}

func (h *Hub) GetUpgrader() *websocket.Upgrader {
	return h.upgrader
}

func (h *Hub) sendToClient(c *Client, data []byte) {
	select {
	case c.sendChan <- data:
	default:
		h.removeClient(c)
	}
}

func (h *Hub) addClient(c *Client) {
	if _, ok := h.clients[c]; !ok {
		h.clients[c] = true
	}
}

func (h *Hub) removeClient(c *Client) {
	if _, ok := h.clients[c]; ok {
		delete(h.clients, c)
	}
}

func (h *Hub) RegisterClient(conn *websocket.Conn) {
	log.Debug("Incoming client connection", "address", conn.RemoteAddr().String())
	NewClient(h, conn)
}

func (h *Hub) Broadcast(msg models.MessageInterface) error {
	data, err := h.serializeMessage(msg)
	if err != nil {
		return err
	}

	h.broadcast <- &PublicMsg{
		//Temp
		channel: string(msg.GetEvent()),
		data:    data,
	}

	return nil
}

func (h *Hub) serializeMessage(msg models.MessageInterface) ([]byte, error) {
	bm := models.BasicMessage{
		Event: msg.GetEvent(),
		Data:  msg,
	}

	return json.Marshal(bm)
}
