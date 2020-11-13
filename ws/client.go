package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/ws/models"
	"github.com/gorilla/websocket"
	"github.com/mailru/easygo/netpoll"
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

		// List of channels which user subscribes to
		subscriptions     map[string]bool
		subscriptionsLock sync.RWMutex

		ctx       context.Context
		cancel    context.CancelFunc
		safeClose sync.Once

		desc *netpoll.Desc
	}
)

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	// Set connection options
	conn.SetReadLimit(maxMessageSize)
	conn.SetWriteDeadline(time.Time{})
	conn.SetReadDeadline(time.Time{})

	ctx, cancel := context.WithCancel(hub.ctx)

	c := &Client{
		hub:           hub,
		conn:          conn,
		sendChan:      make(chan []byte, 256),
		subscriptions: make(map[string]bool),

		ctx:    ctx,
		cancel: cancel,
	}

	var err error
	c.desc, err = netpoll.HandleRead(conn.UnderlyingConn())
	if err != nil {
		//TODO Check Error
		log.Error(err)
	}

	c.register()

	// Allow collection of memory referenced by the caller by doing all work in new goroutines.
	go c.writePump()

	// Send hello message
	c.send(&models.SystemMessage{Message: models.SysMessageHello})

	return c
}

func (cl *Client) register() {
	cl.hub.registerChan <- cl
}

func (cl *Client) unregister() {
	cl.hub.unregisterChan <- cl
}

// readPump pumps messages from the websocket connection to the hub.
func (cl *Client) readPump() {
	defer cl.unregister()

	for {
		_, msgBytes, err := cl.conn.ReadMessage()
		if err != nil {
			log.Debugf("Ws error: %s", err.Error())
			return
		}

		message, err := cl.parseMessage(msgBytes)
		if err != nil {
			// Send user unknown message command
			cl.send(&models.SystemMessage{Message: models.SysMessageUnknownCommand})
			continue
		}

		err = cl.handleMessage(message)
		if err != nil {
			log.Errorf("WS handleMessage: %s", err.Error())
			continue
		}
	}
}

func (cl *Client) Reader() {
	_, msgBytes, err := cl.conn.ReadMessage()
	if err != nil {
		log.Debugf("Ws error: %s", err.Error())
		cl.unregister()
		//TODO check cancel on reader
		//cl.cancel()
		return
	}

	message, err := cl.parseMessage(msgBytes)
	if err != nil {
		// Send user unknown message command
		cl.send(&models.SystemMessage{Message: models.SysMessageUnknownCommand})
	}

	err = cl.handleMessage(message)
	if err != nil {
		log.Errorf("WS handleMessage: %s", err.Error())
		return
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (cl *Client) writePump() {
	defer cl.Close()

	for {
		select {
		case message, ok := <-cl.sendChan:
			if !ok {
				// The hub closed the channel.
				cl.conn.WriteMessage(websocket.CloseMessage, []byte{})
				log.Debug("Closed")
				return
			}

			w, err := cl.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Debug("NextWriter Error")
				return
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-time.After(pingInterval):
			log.Debug("Ping message")
			if err := cl.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Debug("Ping errror")
				return
			}
		case <-cl.ctx.Done():
			log.Debug("Ctx.Done")
			return
		}
	}
}

func (cl *Client) isSubscribed(channel string) bool {
	if channel == "" {
		return false
	}

	cl.subscriptionsLock.RLock()
	defer cl.subscriptionsLock.RUnlock()

	_, ok := cl.subscriptions[channel]
	return ok
}

func (cl *Client) parseMessage(message []byte) (models.ClientMessage, error) {
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

func (cl *Client) handleMessage(msg models.ClientMessage) error {
	var err error

	switch m := msg.(type) {
	case *models.ClientMessageSubscribe:
		cl.subscribe(m.Channels...)
	case *models.ClientMessageUnsubscribe:
		cl.unsubscribe(m.Channels...)
	default:
		return fmt.Errorf("Unknown client message type")
	}

	return err
}

func (cl *Client) send(msg models.MessageInterface) error {
	data, err := cl.hub.serializeMessage(msg)
	if err != nil {
		return err
	}

	cl.hub.sendToClient(cl, data)

	return nil
}

func (cl *Client) subscribe(channels ...string) {
	cl.subscriptionsLock.Lock()
	defer cl.subscriptionsLock.Unlock()

	for _, c := range channels {
		if c != "" {
			c = strings.ToLower(c)
			cl.subscriptions[c] = true
		}
	}

	cl.send(&models.SystemMessage{Message: models.SysMessageSubscribed, Description: strings.Join(channels, ", ")})
}

func (cl *Client) unsubscribe(channels ...string) {
	cl.subscriptionsLock.Lock()
	defer cl.subscriptionsLock.Unlock()

	for _, c := range channels {
		if c != "" {
			c = strings.ToLower(c)
			delete(cl.subscriptions, c)
		}
	}

	cl.send(&models.SystemMessage{Message: models.SysMessageUnsubscribed, Description: strings.Join(channels, ", ")})
}

// Close will close send channel and connection once
func (cl *Client) Close() {
	cl.safeClose.Do(func() {
		cl.conn.Close()
		close(cl.sendChan)
	})
}
