package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Create a global map to store active WebSocket clients
var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex // To synchronize access to the clients map

// WebSocketHandler handles WebSocket connections
func WebSocketHandler(c *gin.Context) {
	// Upgrade HTTP connection to WebSocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // Allow any origin for simplicity
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open websocket connection"})
		return
	}
	defer conn.Close()

	// Add the new connection to the clients map
	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	// Handle messages from the WebSocket connection
	for {
		// Set read deadline
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		// Read the message from the WebSocket connection
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// If there's an error (e.g., client disconnects), remove the client
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			return
		}

		// Broadcast the message to all other connected clients
		broadcastMessage(msg)
	}
}

// broadcastMessage broadcasts the message to all connected Websocket clients
func broadcastMessage(msg []byte) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		// Skip sending the message to the sender
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
