#!/bin/bash

# --- UPDATE & INSTALL ---
sudo apt update && sudo apt upgrade -y
sudo apt install -y nginx certbot python3-certbot-nginx golang git ufw

# --- CONFIGURE FIREWALL ---
sudo ufw allow OpenSSH
sudo ufw allow 'Nginx Full'
sudo ufw --force enable

# --- SETUP GO ENVIRONMENT ---
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc

# --- CREATE WEBSOCKET SERVER ---
mkdir -p ~/websocket-server && cd ~/websocket-server
cat <<EOF > main.go
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	log.Println("WebSocket server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
EOF

# --- INIT GO MODULE ---
go mod init websocket-server
go get github.com/gorilla/websocket
go build -o ws-server main.go

# --- SETUP SYSTEMD SERVICE ---
sudo tee /etc/systemd/system/ws-server.service > /dev/null <<EOF
[Unit]
Description=Golang WebSocket Server
After=network.target

[Service]
ExecStart=/home/ubuntu/websocket-server/ws-server
WorkingDirectory=/home/ubuntu/websocket-server
Restart=always
User=ubuntu
Environment=GOPATH=/home/ubuntu/go

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl enable ws-server
sudo systemctl start ws-server

# --- SETUP NGINX PROXY ---
read -p "Enter your domain name (e.g., ws.example.com): " DOMAIN
sudo tee /etc/nginx/sites-available/ws > /dev/null <<EOF
server {
    listen 80;
    server_name $DOMAIN;

    location /ws/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host \$host;
    }
}
EOF

sudo ln -s /etc/nginx/sites-available/ws /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx

# --- SSL WITH CERTBOT ---
sudo certbot --nginx -d $DOMAIN

echo "✅ WebSocket server is live at wss://$DOMAIN/ws/"
