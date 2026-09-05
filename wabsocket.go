package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)

var clientsMutex sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	println("WebSocket client connected")

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	clientsMutex.Lock()

	clients[conn] = true

	clientsMutex.Unlock()

	conn.WriteJSON(
		GetStatus(),
	)

	defer func() {

		clientsMutex.Lock()

		delete(
			clients,
			conn,
		)

		clientsMutex.Unlock()

		conn.Close()

	}()

	for {

		_, _, err := conn.ReadMessage()

		if err != nil {
			return
		}

	}

}

func broadcastStatus(status DemonStatus) {

	clientsMutex.Lock()

	defer clientsMutex.Unlock()

	for client := range clients {

		err := client.WriteJSON(status)

		if err != nil {

			client.Close()

			delete(
				clients,
				client,
			)

		}

	}

}
