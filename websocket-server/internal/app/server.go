package app

import (
	"WS/websocket-server/internal/adapter"
	"WS/websocket-server/internal/service"
	"net/http"
)

type Server struct {
	Handler *adapter.Handler
}

func NewServer() *Server {
	svc := service.NewWebsoquetService()
	handler := adapter.NewHandler(svc)
	return &Server{
		Handler: handler,
	}
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/ws", s.Handler.ServeWS)
	http.HandleFunc("/internal/send", s.Handler.HandleInternalSend)
	return http.ListenAndServe(addr, nil)
}
