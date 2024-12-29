package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins; adjust for production
	},
}

// WebSocketHandler handles WebSocket connections
func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection established")

	// Handle messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket error:", err)
			break
		}
		log.Printf("Received: %s\n", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
