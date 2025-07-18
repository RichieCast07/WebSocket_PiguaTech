package domain

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

type Client interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(messageType int, msg []byte) error
	Close() error
}