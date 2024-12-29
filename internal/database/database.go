package database

import (
	"fmt"
	"log"
	"realtime-doc-editor-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBConnection initialzes a new PostgreSQL connection using GORM
func NewDBConnection(cfg *config.Config) (*gorm.DB, error) {
	// GORM connection string for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	// Open the connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database using GORM")
	return db, nil
}
