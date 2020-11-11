package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/ws/models"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

const (
	// Maximum message size allowed from peer.
	maxMessageSize = 512

	// Send pings to peer with this period. Must be less than pongWait.
	pingInterval = 30 * time.Second
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type (
	UserMsg struct {
		Type    string          `json:"type"`
		Payload json.RawMessage `json:"payload"`
	}

	Client struct {
		hub *Hub

		// The websocket connection.
		conn *websocket.Conn

		// Buffered channel of outbound messages.
		sendChan chan []byte

		payload string

		// List of channels which user subscribes to
		subscriptions     map[string]bool
		subscriptionsLock sync.RWMutex

		safeClose sync.Once
	}
)

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	// Set connection options
	conn.SetReadLimit(maxMessageSize)
	conn.SetWriteDeadline(time.Time{})
	conn.SetReadDeadline(time.Time{})

	c := &Client{
		hub:           hub,
		conn:          conn,
		sendChan:      make(chan []byte, 256),
		subscriptions: make(map[string]bool),
		safeClose:     sync.Once{},
	}

	c.register()

	// Allow collection of memory referenced by the caller by doing all work in new goroutines.
	go c.writePump()
	go c.readPump()

	// Send hello message
	c.send(&models.SystemMessage{Message: models.SysMessageHello})

	return c
}

func (this *Client) register() {
	this.hub.registerChan <- this
}

func (this *Client) unregister() {
	this.hub.unregisterChan <- this
}

// readPump pumps messages from the websocket connection to the hub.
func (this *Client) readPump() {
	defer this.unregister()

	for {
		_, msgBytes, err := this.conn.ReadMessage()
		if err != nil {
			log.Debugf("Ws error: %s", err.Error())
			break
		}

		message, err := this.parseMessage(msgBytes)
		if err != nil {
			// Send user unknown message command
			this.send(&models.SystemMessage{Message: models.SysMessageUnknownCommand})
			continue
		}

		err = this.handleMessage(message)
		if err != nil {
			log.Errorf("WS handleMessage: %s", err.Error())
			continue
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (this *Client) writePump() {
	defer this.Close()

	for {
		select {
		case message, ok := <-this.sendChan:
			if !ok {
				// The hub closed the channel.
				this.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := this.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-time.After(pingInterval):
			if err := this.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (this *Client) isSubscribed(channel string) bool {
	if channel == "" {
		return false
	}

	this.subscriptionsLock.RLock()
	defer this.subscriptionsLock.RUnlock()

	_, ok := this.subscriptions[channel]
	return ok
}

func (this *Client) parseMessage(message []byte) (models.ClientMessage, error) {
	// Trim empty bytes if any
	msgBytes := bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

	// Decode json
	var um UserMsg
	err := json.Unmarshal(msgBytes, &um)
	if err != nil {
		return nil, err
	}

	var msg models.ClientMessage
	switch um.Type {
	case models.ClientMessageTypeSubscribe:
		msg = &models.ClientMessageSubscribe{}
	case models.ClientMessageTypeUnsubscribe:
		msg = &models.ClientMessageUnsubscribe{}
	default:
		return nil, fmt.Errorf("Unknown message type")
	}

	err = json.Unmarshal(um.Payload, &msg)
	return msg, err
}

func (this *Client) handleMessage(msg models.ClientMessage) error {
	var err error

	switch m := msg.(type) {
	case *models.ClientMessageSubscribe:
		this.subscribe(m.Channels...)
	case *models.ClientMessageUnsubscribe:
		this.unsubscribe(m.Channels...)
	default:
		return fmt.Errorf("Unknown client message type")
	}

	return err
}

func (this *Client) send(msg models.MessageInterface) error {
	data, err := this.hub.serializeMessage(msg)
	if err != nil {
		return err
	}

	this.hub.sendToClient(this, data)

	return nil
}

func (this *Client) subscribe(channels ...string) {
	this.subscriptionsLock.Lock()
	defer this.subscriptionsLock.Unlock()

	for _, c := range channels {
		if c != "" {
			c = strings.ToLower(c)
			this.subscriptions[c] = true
		}
	}

	this.send(&models.SystemMessage{Message: models.SysMessageSubscribed, Description: strings.Join(channels, ", ")})
}

func (this *Client) unsubscribe(channels ...string) {
	this.subscriptionsLock.Lock()
	defer this.subscriptionsLock.Unlock()

	for _, c := range channels {
		if c != "" {
			c = strings.ToLower(c)
			delete(this.subscriptions, c)
		}
	}

	this.send(&models.SystemMessage{Message: models.SysMessageUnsubscribed, Description: strings.Join(channels, ", ")})
}

// Close will close send channel and connection once
func (this *Client) Close() {
	this.safeClose.Do(func() {
		this.conn.Close()
		close(this.sendChan)
	})
}
