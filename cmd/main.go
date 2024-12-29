package main

import (
	"log"
	"realtime-doc-editor-backend/config"
	"realtime-doc-editor-backend/internal/database"
	"realtime-doc-editor-backend/internal/models"
	"realtime-doc-editor-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configurations
	cfg := config.LoadConfig()

	// Initialize the database connection
	db, err := database.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations (this will create the documents table)
	err = db.AutoMigrate(models.Document{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	// Initialize the router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, db, cfg)

	// Start server
	log.Printf("Server running on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
