package battery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"websocket-server/model"
)

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

		var data model.BatteryInfo
		if err := json.Unmarshal(msg, &data); err != nil {
			log.Println("Invalid JSON:", err)
			jsonResp, _ := model.ResponseError("Invalid JSON:" + err.Error())
			conn.WriteMessage(websocket.TextMessage, jsonResp)
			continue
		}

		log.Printf("Battery info received: %+v\n", data)

		// Return success
		jsonResp, _ := model.ResponseSuccess(data, model.TypeBattInfo)
		conn.WriteMessage(websocket.TextMessage, jsonResp)
	}
}
