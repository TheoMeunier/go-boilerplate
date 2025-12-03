package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}

	if cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBName == "" {
		log.Fatal("missing required environment variables for database")
	}

	return cfg
}
