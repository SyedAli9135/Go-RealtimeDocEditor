package unit_tests

import (
	"os"
	"realtime-doc-editor-backend/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("SERVER_ADDRESS", ":9000")
	os.Setenv("DB_USER", "testuser")
	cfg := config.LoadConfig()

	assert.Equal(t, ":9000", cfg.ServerAddress)
	assert.Equal(t, "testuser", cfg.DBUser)
}
