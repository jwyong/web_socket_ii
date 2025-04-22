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

	fmt.Println("WebSocket server listening on http://0.0.0.0" + config.TLSPort)

	// TODO: JAY_LOG - local test
	// CertFile := "certs/server.crt"
	// KeyFile := "certs/server.key"

	// Connect with TSL
	if err := http.ListenAndServeTLS(config.TLSPort, config.CertFile, config.KeyFile, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
