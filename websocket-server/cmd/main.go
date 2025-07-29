package main

import (
	"WS/websocket-server/internal/app"
	"fmt"
	"log"
)

func main() {
	server := app.NewServer()
	addr := ":8080"
	fmt.Printf("Servidor WebSocket ejecutándose en %s\n", addr)
	if err := server.Start(addr); err != nil {
		log.Fatal("Error fatal en el servidor:", err)
	}
}
