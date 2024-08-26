package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}


type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}


type ServerConfig struct {
	Port string
}


func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			DBName:   getEnv("DB_NAME"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT"),
		},
	}

	return cfg, nil
}


func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
