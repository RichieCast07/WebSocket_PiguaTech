package main

import (
	"fmt"
	"log"
	"WS/websocket-server/internal/app"
)

func main() {
	server := app.NewServer()
	addr := ":3000"
	fmt.Printf("Servidor WebSocket ejecutándose en %s\n", addr)
	if err := server.Start(addr); err != nil {
		log.Fatal("Error fatal en el servidor:", err)
	}
}