package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"WS/websocket-server/internal/domain"
	"WS/websocket-server/internal/service"

	"github.com/gorilla/websocket"
)

type Handler struct {
	Service *service.WebsoquetService
}

func NewHandler(svc *service.WebsoquetService) *Handler {
	return &Handler{Service: svc}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar a WebSocket:", err)
		return
	}
	client := NewClient(conn)

	accountID := r.URL.Query().Get("account")
	if accountID == "" {
		log.Println("Error: No se proporcionó el accountID en la URL de conexión.")
		client.Close()
		return
	}

	h.Service.RegisterClient(accountID, client)
	go h.handleClientMessages(accountID, client)
}

func (h *Handler) handleClientMessages(accountID string, client domain.Client) {
	defer func() {
		h.Service.RemoveClient(accountID)
	}()
	for {
		messageType, msg, err := client.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error inesperado leyendo mensaje del cliente %s: %v", accountID, err)
			}
			break
		}

		var message domain.Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error decodificando JSON del cliente:", err)
			continue
		}

		h.Service.SendMessageToAccount(message.Receiver, messageType, msg)
	}
}

type InternalSendRequest struct {
	Receiver string          `json:"receiver"`
	Payload  json.RawMessage `json:"payload"`
}

func (h *Handler) HandleInternalSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req InternalSendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}

	if req.Receiver == "" {
		http.Error(w, "'receiver' es un campo requerido", http.StatusBadRequest)
		return
	}
	if len(req.Payload) == 0 {
		http.Error(w, "'payload' es un campo requerido", http.StatusBadRequest)
		return
	}

	h.Service.SendMessageToAccount(req.Receiver, websocket.TextMessage, []byte(req.Payload))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Mensaje encolado para entrega.")
}