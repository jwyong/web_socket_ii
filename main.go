package main

import (
	"fmt"
	"log"
	"net/http"

	"websocket-server/battery"
	// "websocket-server/config"
)

func main() {
	http.HandleFunc("/ws/battery", battery.HandleWebSocket)

	// TODO: JAY_LOG - local test
	// CertFile := "certs/server.crt"
	// KeyFile := "certs/server.key"

	fmt.Println("WebSocket server listening on http://0.0.0.0:8080")

	// Connect without TSL
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	// Connect with TSL
	// if err := http.ListenAndServeTLS(config.TLSPort, config.CertFile, config.KeyFile, nil); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }
}
