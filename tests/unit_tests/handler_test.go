package unit_tests

import (
	"net/http/httptest"
	"net/url"
	"realtime-doc-editor-backend/internal/handlers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocketHandler(t *testing.T) {
	// Create a test Gin engine
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/ws", handlers.WebSocketHandler)

	// Start a test HTTP server with the router
	server := httptest.NewServer(router)
	defer server.Close()

	// Parse the server's URL and establish a WebSocket connection
	serverURL := url.URL{Scheme: "ws", Host: server.Listener.Addr().String(), Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	assert.NoError(t, err, "Failed to establish WebSocket connection")
	defer conn.Close()

	// Send a test message
	testMessage := "hello"
	err = conn.WriteMessage(websocket.TextMessage, []byte(testMessage))
	assert.NoError(t, err, "Failed to send message")

	// Read the echoed message
	messageType, receivedMessage, err := conn.ReadMessage()
	assert.NoError(t, err, "Failed to read message")
	assert.Equal(t, websocket.TextMessage, messageType, "Unexpected message type")
	assert.Equal(t, testMessage, string(receivedMessage), "Unexpected message content")
}
