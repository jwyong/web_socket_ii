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

	fmt.Println("WebSocket server listening on https://0.0.0.0" + config.TLSPort)

	if err := http.ListenAndServeTLS(config.TLSPort, config.CertFile, config.KeyFile, nil); err != nil {
		log.Fatal("ListenAndServeTLS:", err)
	}
}
