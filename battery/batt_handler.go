package battery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type BatteryInfo struct {
	Level       int  `json:"level"`
	Charging    bool `json:"charging"`
	Temperature int  `json:"temperature"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var data BatteryInfo
		if err := json.Unmarshal(msg, &data); err != nil {
			log.Println("Invalid JSON:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid JSON"))
			continue
		}

		log.Printf("Battery info received: %+v\n", data)
		message := fmt.Sprintf("Battery info received: %+v", data)

		// Return ack on receiving battery info.
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
