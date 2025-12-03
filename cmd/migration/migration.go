package main

import (
	"boilerplate/internal/infra/config"
	"boilerplate/internal/infra/db"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewDatabase(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.Close()

	files, err := filepath.Glob("./internal/infra/db/migrations/*.sql")
	if err != nil {
		log.Fatalf("Error reading migrations: %v", err)
	}

	for _, file := range files {
		fmt.Println("Running migration: ", file)

		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading migration file: %v", err)
		}

		_, err = database.Conn.Exec(string(content))
		if err != nil {
			log.Fatalf("Error running migration: %v", err)
		}
	}

	fmt.Println("Migrations completed successfully")
}
