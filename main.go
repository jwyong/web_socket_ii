package main

import (
	"fmt"
	"log"
	"net/http"

	"websocket-server/battery"
	"websocket-server/config"
)

func main() {
	http.HandleFunc("/ws/battery", battery.HandleWebSocket)

	port := ":8080"
	fmt.Println("WebSocket server listening on http://0.0.0.0" + port)

	// TODO: JAY_LOG - no SSL for now
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
