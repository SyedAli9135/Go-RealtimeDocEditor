package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config struct for application configurations
type Config struct {
	ServerAddress string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
}

// LoadConfig loads configuration from environment variables of defaults
func LoadConfig() *Config {
	viper.SetDefault("SERVER_ADDRESS", ":8000")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5433")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "realtime_docs")
	viper.SetDefault("DB_SSL_MODE", "disable")

	viper.AutomaticEnv()

	config := &Config{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetString("DB_PORT"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
		DBSSLMode:     viper.GetString("DB_SSL_MODE"),
	}

	log.Println("Configuration loaded successfully")
	return config
}
