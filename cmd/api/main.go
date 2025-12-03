package main

import (
	"boilerplate/internal/infra/config"
	"boilerplate/internal/infra/db"
	"boilerplate/internal/infra/http"
	"boilerplate/internal/infra/storage"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewDatabase(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	database.Close()

	_, err = storage.InitStorage()

	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	http.NewHandler()
}
