package middleware

import (
	"net/http"
	"realtime-doc-editor-backend/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware checks the validity of the JWT token in the request header
func JWTMiddleware(c *gin.Context) {
	// Get the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	// Split the header value to extract the token
	tokenString := strings.Split(authHeader, "Bearer ")[1]
	token, err := auth.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// Store the parsed token in the context for further use
	c.Set("token", token)
	c.Next()
}
