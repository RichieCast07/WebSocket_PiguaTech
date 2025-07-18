package adapter

import (
	"WS/websocket-server/internal/domain"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Conn: conn}
}

func (c *Client) ReadMessage() (int, []byte, error) {
	return c.Conn.ReadMessage()
}

func (c *Client) WriteMessage(messageType int, msg []byte) error {
	return c.Conn.WriteMessage(messageType, msg)
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

var _ domain.Client = (*Client)(nil)