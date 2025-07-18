package service

import (
	"log"
	"sync"
	"WS/websocket-server/internal/domain"
)

type WebsoquetService struct {
	Clients map[string]domain.Client
	mu      sync.Mutex
}

func NewWebsoquetService() *WebsoquetService {
	return &WebsoquetService{
		Clients: make(map[string]domain.Client),
	}
}

func (s *WebsoquetService) RegisterClient(accountID string, client domain.Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Clients[accountID] = client
	log.Printf("Cliente registrado: %s\n", accountID)
}

func (s *WebsoquetService) SendMessageToAccount(receiver string, messageType int, msg []byte) {
	s.mu.Lock()
	client, ok := s.Clients[receiver]
	s.mu.Unlock()

	if !ok {
		log.Printf("Cliente %s no encontrado o desconectado\n", receiver)
		return
	}

	err := client.WriteMessage(messageType, msg)
	if err != nil {
		log.Printf("Error enviando mensaje a %s: %v. Eliminando cliente.", receiver, err)
		s.RemoveClient(receiver)
	}
}

func (s *WebsoquetService) RemoveClient(accountID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	client, ok := s.Clients[accountID]
	if ok {
		client.Close()
		delete(s.Clients, accountID)
		log.Printf("Cliente desconectado: %s\n", accountID)
	}
}