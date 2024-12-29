package auth

import (
	"fmt"
	"os"
	"realtime-doc-editor-backend/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT Secret Key for signing tokens
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// CreateToken generates a JWT token for a given user
func CreateToken(user models.User) (string, error) {
	// Set the claims for the token (Expiration, Subject, etc.)
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // 1 day expiration
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString string) (*jwt.Token, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
