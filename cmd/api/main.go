package main

import (
	"boilerplate/internal/infra/config"
	"boilerplate/internal/infra/db"
	"boilerplate/internal/infra/http"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewDatabase(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	database.Close()

	http.NewHandler()
}
