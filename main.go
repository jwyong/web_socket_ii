package main

import (
	"fmt"
	"log"
	"net/http"

	"websocket-server/battery"
)

func main() {
	http.HandleFunc("/ws/battery", battery.HandleWebSocket)

	port := ":8080"
	fmt.Println("WebSocket server listening on", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
